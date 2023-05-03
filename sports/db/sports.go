package db

import (
	"database/sql"
	"git.neds.sh/matty/entain/sports/proto/sports"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"
	"strings"
	"sync"
	"time"
)

// SportsRepo provides repository access to sports.
type SportsRepo interface {
	// Init will initialise our sports repository.
	Init() error

	// List will return a list of sports.
	List(filter *sports.ListSportsRequestFilter, sort *sports.ListSportsRequestSort) ([]*sports.Sport, error)
	Get(request *sports.GetSportRequest) (*sports.Sport, error)
}

type sportsRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewSportsRepo creates a new sports repository.
func NewSportsRepo(db *sql.DB) SportsRepo {
	return &sportsRepo{db: db}
}

// Init prepares the sports repository dummy data.
func (r *sportsRepo) Init() error {
	var err error

	r.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy sports.
		err = r.seed()
	})

	return err
}

func (r *sportsRepo) Get(request *sports.GetSportRequest) (*sports.Sport, error) {
	var (
		query string
		args  []interface{}
	)

	query = getSportsQueries()[getEvent]

	args = append(args, request.Id)
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	sportsList, err := r.scanSports(rows)
	if err != nil {
		return nil, err
	} else if len(sportsList) == 0 {
		// No sports found
		return nil, sql.ErrNoRows
	} else {
		// return sport.
		return sportsList[0], nil
	}
}

func (r *sportsRepo) List(filter *sports.ListSportsRequestFilter, sort *sports.ListSportsRequestSort) ([]*sports.Sport, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getSportsQueries()[eventsList]

	query, args = r.applyFilter(query, filter)

	query = r.applySorting(query, sort)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return r.scanSports(rows)
}

func (r *sportsRepo) applyFilter(query string, filter *sports.ListSportsRequestFilter) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
	)

	if filter == nil {
		return query, args
	}

	if filter.Visible != nil {
		clauses = append(clauses, "visible = ?")
		args = append(args, filter.Visible)
	}

	if len(clauses) != 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}

	return query, args
}

func (r *sportsRepo) applySorting(query string, sort *sports.ListSportsRequestSort) string {
	if sort == nil {
		return query
	}

	if sort.Field == nil {
		return query
	}

	switch *sort.Field.Enum() {
	case sports.ListSportsRequestSort_ADVERTISED_START_TIME:
		query += " ORDER BY advertised_start_time"
	case sports.ListSportsRequestSort_ID:
		query += " ORDER BY id"
	case sports.ListSportsRequestSort_NAME:
		query += " ORDER BY name"
	}

	if sort.Order == nil {
		return query
	}
	switch *sort.Order.Enum() {
	case sports.ListSportsRequestSort_ASC:
		query += " ASC"
	case sports.ListSportsRequestSort_DESC:
		query += " DESC"
	}
	return query
}

func (m *sportsRepo) scanSports(
	rows *sql.Rows,
) ([]*sports.Sport, error) {
	var sports_list []*sports.Sport

	for rows.Next() {
		var sport sports.Sport
		var advertisedStart time.Time

		if err := rows.Scan(&sport.Id, &sport.Name, &sport.Visible, &advertisedStart); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		sport.AdvertisedStartTime = ts

		sports_list = append(sports_list, &sport)
	}

	return sports_list, nil
}

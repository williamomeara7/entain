package service

import (
	"git.neds.sh/matty/entain/sports/db"
	"git.neds.sh/matty/entain/sports/proto/sports"
	"golang.org/x/net/context"
)

type Sports interface {
	// ListSports will return a collection of Sports.
	ListSports(ctx context.Context, in *sports.ListSportsRequest) (*sports.ListSportsResponse, error)
	GetSport(ctx context.Context, in *sports.GetSportRequest) (*sports.Sport, error)
}

// racingService implements the Racing interface.
type sportsService struct {
	sportsRepo db.SportsRepo
}

// NewSportService instantiates and returns a new sportsService.
func NewSportService(sportsRepo db.SportsRepo) Sports {
	return &sportsService{sportsRepo}
}

func (s *sportsService) ListSports(ctx context.Context, in *sports.ListSportsRequest) (*sports.ListSportsResponse, error) {
	sportsList, err := s.sportsRepo.List(in.Filter, in.Sort)
	if err != nil {
		return nil, err
	}

	return &sports.ListSportsResponse{Sports: sportsList}, nil
}

func (s *sportsService) GetSport(ctx context.Context, in *sports.GetSportRequest) (*sports.Sport, error) {
	sport, err := s.sportsRepo.Get(in)
	if err != nil {
		return nil, err
	}
	return sport, nil
}

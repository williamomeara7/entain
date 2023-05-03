package db

const (
	eventsList = "list"
	getEvent   = "get"
)

func getSportsQueries() map[string]string {
	return map[string]string{
		eventsList: `
			SELECT 
				id, 
				name, 
				visible, 
				advertised_start_time 
			FROM events
		`,
		getEvent: `
			SELECT
				id,
				name,
				visible,	
				advertised_start_time
			FROM events
			WHERE id = ?
		`,
	}
}

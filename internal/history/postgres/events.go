package postgres

const (
	eventTablename  = "event"
	userTablename   = "user"
	sharetTablename = "event_user"
)

// func (p *PostgresHistory) Create(eventId, title, userId string, isDeleted bool) (*responses.HistoryResponse, error){
// query := fmt.Sprintf("INSERT INTO %s (id, title, user_id, is_deleted) VALUES($1, $2, $3, $4)")

// }

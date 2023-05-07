package postgres

import (
	"fmt"

	"github.com/Str1kez/SportiqSubscriptionService/internal/responses"
	log "github.com/sirupsen/logrus"
)

const (
	eventTablename  = "event"
	userTablename   = "user"
	sharedTablename = "event_user"
)

func (p *PostgresHistory) Create(eventId, title, userId string, isDeleted bool) (*responses.HistoryResponse, error) {
	queryEvent := fmt.Sprintf(`INSERT INTO %s (id, title, is_deleted) VALUES($1, $2, $3)`, eventTablename)
	queryUser := fmt.Sprintf(`INSERT INTO "%s" (id) VALUES($1)`, userTablename)
	queryShared := fmt.Sprintf(`INSERT INTO %s (event_id, user_id) VALUES($1, $2)`, sharedTablename)

	tx, err := p.connection.Beginx()
	if err != nil {
		log.Error("Can't begin transaction", err)
		return nil, err
	}
	row := tx.QueryRowx(queryEvent, eventId, title, isDeleted)
	if err = row.Err(); err != nil {
		log.Errorf("Can't insert row: %v\n", err)
		tx.Rollback()
		return nil, err
	}
	row = tx.QueryRowx(queryUser, userId)
	if err = row.Err(); err != nil {
		log.Errorf("Can't insert row: %v\n", err)
		tx.Rollback()
		return nil, err
	}
	row = tx.QueryRowx(queryShared, eventId, userId)
	if err = row.Err(); err != nil {
		log.Errorf("Can't insert row: %v\n", err)
		tx.Rollback()
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		log.Errorf("Couldn't commit transaction", err)
		tx.Rollback()
		return nil, err
	}

	return &responses.HistoryResponse{EventId: eventId, UserId: userId, EventTitle: title, IsDeleted: isDeleted}, nil
}

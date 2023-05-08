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

func (p *PostgresHistory) Create(eventId string, title interface{}, usersId []string, isDeleted bool) error {
	queryEvent := fmt.Sprintf(`INSERT INTO %s (id, title, is_deleted) VALUES($1, $2, $3)`, eventTablename)
	queryUser := fmt.Sprintf(`INSERT INTO "%s" (id) VALUES($1) ON CONFLICT DO NOTHING`, userTablename)
	queryShared := fmt.Sprintf(`INSERT INTO %s (event_id, user_id) VALUES($1, $2)`, sharedTablename)

	tx, err := p.connection.Beginx()
	if err != nil {
		log.Error("Can't begin transaction", err)
		return err
	}
	row := tx.QueryRowx(queryEvent, eventId, title, isDeleted)
	if err = row.Err(); err != nil {
		log.Errorf("Can't insert row: %v\n", err)
		tx.Rollback()
		return err
	}
  for _, u := range usersId {
    row = tx.QueryRowx(queryUser, u)
    if err = row.Err(); err != nil {
      log.Errorf("Can't insert row: %v\n", err)
      tx.Rollback()
      return err
    }
    row = tx.QueryRowx(queryShared, eventId, u)
    if err = row.Err(); err != nil {
      log.Errorf("Can't insert row: %v\n", err)
      tx.Rollback()
      return err
    }
  }
	if err = tx.Commit(); err != nil {
		log.Errorf("Couldn't commit transaction", err)
		tx.Rollback()
		return err
	}

	return nil
}

func (p *PostgresHistory) Get(userId string) ([]*responses.HistoryResponse, error) {
	row_query := `SELECT %s.id AS id, title, is_deleted 
                FROM %s 
                JOIN %s ON %s.event_id = %s.id 
                WHERE %s.user_id = $1 
                ORDER BY %s.created_at DESC;`
	query := fmt.Sprintf(row_query, eventTablename, eventTablename, sharedTablename,
		sharedTablename, eventTablename, sharedTablename, eventTablename)
	rows, err := p.connection.Queryx(query, userId)
	if err != nil {
		log.Errorf("Couldn't get rows: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	// ! Неоптимально, лучше сделать через транзакцию с проверкой на кол-во строк
	response := make([]*responses.HistoryResponse, 0, 100)
	for rows.Next() {
		var temp responses.HistoryResponse
		if err := rows.StructScan(&temp); err != nil {
			log.Errorf("Can't parse row: %v\n", err)
			return nil, err
		}
		response = append(response, &temp)
	}
	return response, nil
}

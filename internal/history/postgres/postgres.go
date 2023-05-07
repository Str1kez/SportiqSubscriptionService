package postgres

import (
	"fmt"

	"github.com/Str1kez/SportiqSubscriptionService/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type PostgresHistory struct {
	connection *sqlx.DB
}

func NewPostgresHistory(config *config.HistoryDBConfig) *PostgresHistory {
	conn, err := sqlx.Connect("pgx", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName))
	if err != nil {
		log.Fatalf("Could't connect to DB: %v\n", err)
		return nil
	}
	return &PostgresHistory{connection: conn}
}

func (p *PostgresHistory) Close() error {
	return p.connection.Close()
}

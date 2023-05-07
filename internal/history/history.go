package history

import (
	"github.com/Str1kez/SportiqSubscriptionService/internal/config"
	"github.com/Str1kez/SportiqSubscriptionService/internal/history/postgres"
)

type history interface {
	Close() error
}

type HistoryDB struct {
	history
}

func NewHistoryDB(config *config.HistoryDBConfig) *HistoryDB {
	return &HistoryDB{postgres.NewPostgresHistory(config)}
}

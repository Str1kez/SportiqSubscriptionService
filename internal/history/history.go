package history

import (
	"github.com/Str1kez/SportiqSubscriptionService/internal/config"
	"github.com/Str1kez/SportiqSubscriptionService/internal/history/postgres"
	"github.com/Str1kez/SportiqSubscriptionService/internal/responses"
)

type history interface {
	Close() error
	Create(eventId, title, userId string, isDeleted bool) (*responses.HistoryResponse, error)
}

type HistoryDB struct {
	history
}

func NewHistoryDB(config *config.HistoryDBConfig) *HistoryDB {
	return &HistoryDB{postgres.NewPostgresHistory(config)}
}

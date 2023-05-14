package history

import (
	"github.com/Str1kez/SportiqSubscriptionService/internal/config"
	"github.com/Str1kez/SportiqSubscriptionService/internal/dto"
	"github.com/Str1kez/SportiqSubscriptionService/internal/history/postgres"
	log "github.com/sirupsen/logrus"
)

type history interface {
	Create(eventId string, title interface{}, usersId []string, isDeleted bool) error
	Get(userId string, size, page int) ([]*dto.HistoryResponse, int, error)
	Close() error
}

type HistoryDB struct {
	history
}

func InitHistoryDB(config *config.HistoryDBConfig) *HistoryDB {
	return &HistoryDB{postgres.NewPostgresHistory(config)}
}

func InitHistoryDBSlice(config *config.HistoryDBConfig, instanceCount uint8) []*HistoryDB {
	if instanceCount == 0 {
		log.Errorln("Zero instances declared")
	}
	var i uint8
	instanceSlice := make([]*HistoryDB, 0, instanceCount)
	for i = 0; i < instanceCount; i++ {
		instanceSlice = append(instanceSlice, InitHistoryDB(config))
	}
	return instanceSlice
}

func GracefulShutdown(instances ...*HistoryDB) {
	for _, i := range instances {
		err := i.Close()
		if err != nil {
			log.Errorf("Error in HistoryDB.Close(): %v\n", err)
		}
	}
}

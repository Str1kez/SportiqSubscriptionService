package db

import (
	"github.com/Str1kez/SportiqSubscriptionService/internal/config"
	"github.com/Str1kez/SportiqSubscriptionService/internal/db/redisjson"
	log "github.com/sirupsen/logrus"
)

type subscription interface {
	DeleteEvent(eventId string) error
	CreateEvent(eventId, userId string) error
	UpdateEventStatus(eventId, eventStatus string) error
	GetUsers(eventId string) ([]string, error)
	Close() error
}

type SubscriptionDB struct {
	subscription
}

func InitSubscriptionDB(config *config.DBConfig) *SubscriptionDB {
	return &SubscriptionDB{subscription: redisjson.NewReJSONDB(config)}
}

func InitSubscriptionDBSlice(config *config.DBConfig, instanceCount uint8) []*SubscriptionDB {
	if instanceCount == 0 {
		log.Errorln("Zero instances declared")
	}
	var i uint8
	instanceSlice := make([]*SubscriptionDB, 0, instanceCount)
	for i = 0; i < instanceCount; i++ {
		instanceSlice = append(instanceSlice, InitSubscriptionDB(config))
	}
	return instanceSlice
}

func GracefulShutdown(instances ...*SubscriptionDB) {
	for _, i := range instances {
		err := i.Close()
		if err != nil {
			log.Errorf("Error in SubscriptionDB.Close(): %v\n", err)
		}
	}
}

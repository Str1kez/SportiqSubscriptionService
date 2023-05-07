package db

import (
	"github.com/Str1kez/SportiqSubscriptionService/internal/config"
	"github.com/Str1kez/SportiqSubscriptionService/pkg/db/redisjson"
	log "github.com/sirupsen/logrus"
)

type subscription interface {
	Close() error
}

type SubscriptionDB struct {
	subscription
}

func NewSubscriptionDB(config *config.DBConfig) *SubscriptionDB {
	return &SubscriptionDB{subscription: redisjson.NewReJSONDB(config)}
}

func NewSubscriptionDBSlice(config *config.DBConfig, instanceCount uint8) []*SubscriptionDB {
	if instanceCount == 0 {
		log.Errorln("Zero instances declared")
	}
	var i uint8
	instanceSlice := make([]*SubscriptionDB, 0, instanceCount)
	for i = 0; i < instanceCount; i++ {
		instanceSlice = append(instanceSlice, NewSubscriptionDB(config))
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

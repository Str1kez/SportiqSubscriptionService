package redisjson

import (
	"context"

	"github.com/Str1kez/SportiqSubscriptionService/internal/config"
	"github.com/nitishm/go-rejson/v4"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

type ReJSONDB struct {
	client  *redis.Client
	handler *rejson.Handler
	config  *config.DBConfig
}

func NewReJSONDB(config *config.DBConfig) *ReJSONDB {
	redisClient := redis.NewClient(&redis.Options{
		Addr: config.Url,
	})
	rejsonHandler := rejson.NewReJSONHandler()
	rejsonHandler.SetGoRedisClient(redisClient)
	instance := &ReJSONDB{client: redisClient, handler: rejsonHandler, config: config}
	instance.healthCheck()
	return instance
}

func (r *ReJSONDB) healthCheck() {
	status := r.client.Ping(context.Background())
	res, err := status.Result()
	if err != nil {
		log.Fatalf("Couldn't connect to Redis DB: %v\n", err)
	}
	log.Debugf("Response from Ping command: %v\n", res)
}

func (r *ReJSONDB) Close() error {
	if err := r.client.FlushAll(context.Background()).Err(); err != nil {
		return err
	}
	if err := r.client.Close(); err != nil {
		return err
	}
	return nil
}

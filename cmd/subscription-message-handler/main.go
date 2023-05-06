package main

import (
	"github.com/Str1kez/SportiqSubscriptionService/internal/config"
	"github.com/Str1kez/SportiqSubscriptionService/pkg/logger"
	"github.com/nitishm/go-rejson/v4"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

func main() {
	logger.InitLogger()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Couldn't initialize config: %v\n", err)
	}
	// TODO: НАПИСАТЬ ЛОГИКУ ДЛЯ ПЕРЕХВАТА ПАНИК С ЗАКРЫТИЕМ CONSUMER
	log.Infoln("Config has been parsed")
	log.Infoln(cfg.Server.Host)
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	rh := rejson.NewReJSONHandler()
	rh.SetGoRedisClient(rdb)
}

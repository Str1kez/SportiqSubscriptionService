package main

import (
	"github.com/Str1kez/SportiqSubscriptionService/internal/config"
	"github.com/Str1kez/SportiqSubscriptionService/pkg/logger"
	"github.com/Str1kez/SportiqSubscriptionService/pkg/mq"
	"github.com/nitishm/go-rejson/v4"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

func handlePanic(consumers ...*mq.MQConsumer) {
	if r := recover(); r != nil {
		mq.GracefulShutdown(consumers...)
		log.Fatalf("Failed by panic: %v\n", r)
	}
}

func main() {
	logger.InitLogger()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Couldn't initialize config: %v\n", err)
	}
	consumers := mq.InitMQConsumer(cfg.MQ)
	defer handlePanic(consumers...)

	mq.HandleMessages(consumers...)

	log.Debugln("checking availability")
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	rh := rejson.NewReJSONHandler()
	rh.SetGoRedisClient(rdb)
}

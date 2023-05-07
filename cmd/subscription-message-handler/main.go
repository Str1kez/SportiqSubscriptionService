package main

import (
	"github.com/Str1kez/SportiqSubscriptionService/internal/config"
	"github.com/Str1kez/SportiqSubscriptionService/pkg/db"
	"github.com/Str1kez/SportiqSubscriptionService/pkg/logger"
	"github.com/Str1kez/SportiqSubscriptionService/pkg/mq"
	log "github.com/sirupsen/logrus"
)

func handlePanic(consumers []*mq.MQConsumer, subscriptionDBInstances []*db.SubscriptionDB) {
	if r := recover(); r != nil {
		mq.GracefulShutdown(consumers...)
		db.GracefulShutdown(subscriptionDBInstances...)
		log.Fatalf("Failed by panic: %v\n", r)
	}
}

func main() {
	logger.InitLogger()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Couldn't initialize config: %v\n", err)
	}
	subscriptionDBInstances := db.NewSubscriptionDBSlice(cfg.DB, cfg.MQ.ConsumerCount)
	consumers := mq.InitMQConsumerSlice(cfg.MQ, subscriptionDBInstances)
	defer handlePanic(consumers, subscriptionDBInstances)

	mq.HandleMessages(consumers...)

	// TODO: Add graceful shutdown on SIGINT
	log.Debugln("checking availability")
}

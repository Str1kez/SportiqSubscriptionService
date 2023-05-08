package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Str1kez/SportiqSubscriptionService/internal/config"
	"github.com/Str1kez/SportiqSubscriptionService/internal/db"
	"github.com/Str1kez/SportiqSubscriptionService/internal/history"
	"github.com/Str1kez/SportiqSubscriptionService/pkg/logger"
	"github.com/Str1kez/SportiqSubscriptionService/pkg/mq"
	log "github.com/sirupsen/logrus"
)

func shutdownService(consumers []*mq.MQConsumer, subscriptionDBInstances []*db.SubscriptionDB, historyDBInstances []*history.HistoryDB) {
	mq.GracefulShutdown(consumers...)
	db.GracefulShutdown(subscriptionDBInstances...)
	history.GracefulShutdown(historyDBInstances...)
	log.Infoln("Service has been disabled")
}

func handlePanic(consumers []*mq.MQConsumer, subscriptionDBInstances []*db.SubscriptionDB, historyDBInstances []*history.HistoryDB) {
	if r := recover(); r != nil {
		shutdownService(consumers, subscriptionDBInstances, historyDBInstances)
		log.Fatalln("Failed by panic")
	}
}

func main() {
	logger.InitLogger()

	log.Infoln("Service is starting")
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Couldn't initialize config: %v\n", err)
	}
	subscriptionDBInstances := db.InitSubscriptionDBSlice(&cfg.DB, cfg.MQ.ConsumerCount)
	historyDBInstances := history.InitHistoryDBSlice(&cfg.HistoryDB, cfg.MQ.ConsumerCount)
	consumers := mq.InitMQConsumerSlice(&cfg.MQ, subscriptionDBInstances, historyDBInstances)
	defer handlePanic(consumers, subscriptionDBInstances, historyDBInstances)

	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go mq.HandleMessages(consumers...)

	<-terminate
	shutdownService(consumers, subscriptionDBInstances, historyDBInstances)
}

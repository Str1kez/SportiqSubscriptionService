package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Str1kez/SportiqSubscriptionService/api"
	"github.com/Str1kez/SportiqSubscriptionService/internal/config"
	"github.com/Str1kez/SportiqSubscriptionService/internal/db"
	"github.com/Str1kez/SportiqSubscriptionService/internal/history"
	"github.com/Str1kez/SportiqSubscriptionService/pkg/logger"
	"github.com/Str1kez/SportiqSubscriptionService/pkg/server"
	log "github.com/sirupsen/logrus"
)

func gracefulShutdown(server *server.Server, subDB *db.SubscriptionDB, historyDB *history.HistoryDB) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	db.GracefulShutdown(subDB)
	history.GracefulShutdown(historyDB)
	if err := server.Shutdown(ctx); err != nil {
		log.Errorf("Error in shutting down\n%s\n", err)
	}
	log.Infoln("Service has been disabled")
}

func handlePanic(server *server.Server, subDB *db.SubscriptionDB, historyDB *history.HistoryDB) {
	if r := recover(); r != nil {
		gracefulShutdown(server, subDB, historyDB)
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

	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	server := server.Server{}
	historyDB := history.InitHistoryDB(&cfg.HistoryDB)
	subscriptionDB := db.InitSubscriptionDB(&cfg.DB)
	httpServer := api.InitHttpServer(subscriptionDB, historyDB)
	defer handlePanic(&server, subscriptionDB, historyDB)

	go func() {
		if err = server.Run(cfg.Server.Host, cfg.Server.Port, httpServer.InitRouters()); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error on server startup: %s\n", err)
		}
	}()

	<-terminate
	gracefulShutdown(&server, subscriptionDB, historyDB)
}

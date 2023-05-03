package main

import (
	"github.com/Str1kez/SportiqSubscriptionService/internal/config"
	"github.com/Str1kez/SportiqSubscriptionService/pkg/logger"
	log "github.com/sirupsen/logrus"
)

func main() {
	logger.InitLogger()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Couldn't initialize config: %v\n", err)
	}
	log.Infoln("Config has been parsed")
	log.Infoln(cfg.Server.Host)
}

package api

import (
	"github.com/Str1kez/SportiqSubscriptionService/api/controllers"
	"github.com/Str1kez/SportiqSubscriptionService/api/middlewares"
	"github.com/Str1kez/SportiqSubscriptionService/internal/db"
	"github.com/Str1kez/SportiqSubscriptionService/internal/history"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	ginlogrus "github.com/toorop/gin-logrus"
	pagination "github.com/webstradev/gin-pagination"
)

type HttpServer struct {
	controller *controllers.Controller
}

func InitHttpServer(subDB *db.SubscriptionDB, historyDB *history.HistoryDB) *HttpServer {
	return &HttpServer{controller: controllers.InitController(subDB, historyDB)}
}

// @title Sportiq Subscription API
// @version 0.1.0
// @description Subscription microservice for Sportiq Project
// @contact.name Str1kez
// @contact.url https://github.com/Str1kez
// @host localhost:8002
// @BasePath /api/v1/subscription
// @securityDefinitions.apikey UserID
// @in header
// @name User
func (h *HttpServer) InitRouters() *gin.Engine {
	router := gin.New()
	router.HandleMethodNotAllowed = true
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.ExposeHeaders = []string{"x-total-count"}
	router.Use(ginlogrus.Logger(log.StandardLogger()), gin.Recovery(), cors.New(config), middlewares.UserMiddleware())

	apiRouter := router.Group(viper.GetString("prefix"))
	{
		apiRouter.POST("/subscribe/:event_id", h.controller.Subscribe)
		apiRouter.POST("/unsubscribe/:event_id", h.controller.Unsubscribe)
		apiRouter.GET("/subscribers/count/:event_id", h.controller.SubscribersCount)
		apiRouter.GET("/history", pagination.Default(), h.controller.History)
		apiRouter.GET("/subscriptions", h.controller.Subscriptions)
	}

	return router
}

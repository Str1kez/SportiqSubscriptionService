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
)

type HttpServer struct {
	controller *controllers.Controller
}

func InitHttpServer(subDB *db.SubscriptionDB, historyDB *history.HistoryDB) *HttpServer {
	return &HttpServer{controller: controllers.InitController(subDB, historyDB)}
}

func (h *HttpServer) InitRouters() *gin.Engine {
	router := gin.New()
	router.HandleMethodNotAllowed = true
	router.Use(ginlogrus.Logger(log.StandardLogger()), gin.Recovery(), middlewares.UserMiddleware(), cors.Default())

	apiRouter := router.Group(viper.GetString("prefix"))
	{
		apiRouter.POST("/subscribe/:event_id", h.controller.Subscribe)
		apiRouter.POST("/unsubscribe/:event_id", h.controller.Unsubscribe)
		apiRouter.GET("/subscribers/count/:event_id", h.controller.SubscribersCount)
		apiRouter.GET("/history", h.controller.History)
	}

	return router
}

package controllers

import (
	"github.com/Str1kez/SportiqSubscriptionService/api/responses"
	"github.com/Str1kez/SportiqSubscriptionService/internal/db"
	"github.com/Str1kez/SportiqSubscriptionService/internal/history"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	subscriptionDB *db.SubscriptionDB
	historyDB      *history.HistoryDB
}

func InitController(subDB *db.SubscriptionDB, historyDB *history.HistoryDB) *Controller {
	return &Controller{subscriptionDB: subDB, historyDB: historyDB}
}

func (ctl *Controller) response(c *gin.Context, status int, data interface{}) {
	if data == nil {
		c.Status(status)
		return
	}
	c.JSON(status, data)
}

func (ctl *Controller) errorResponse(c *gin.Context, status int, errorType string, err error) {
	response := responses.ErrorResponse{Detail: []responses.ErrorInfo{{Msg: err.Error(), Type: errorType}}}
	c.AbortWithStatusJSON(status, response)
}

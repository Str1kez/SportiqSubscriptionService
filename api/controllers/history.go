package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) History(c *gin.Context) {
	historyResponses, err := ctl.historyDB.Get(c.GetHeader("User"))
	if err != nil {
		ctl.errorResponse(c, http.StatusBadRequest, "subscription.history", err)
		return
	}
	ctl.response(c, http.StatusOK, historyResponses)
}

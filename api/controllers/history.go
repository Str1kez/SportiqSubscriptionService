package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) History(c *gin.Context) {
	page := c.GetInt("page")
	size := c.GetInt("size")
	historyResponses, totalCount, err := ctl.historyDB.Get(c.GetHeader("User"), size, page)
	if err != nil {
		ctl.errorResponse(c, http.StatusBadRequest, "subscription.history", err)
		return
	}
	c.Header("X-Total-Count", strconv.FormatInt(int64(totalCount), 10))
	ctl.response(c, http.StatusOK, historyResponses)
}

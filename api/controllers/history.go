package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary History
// @ID history
// @Description Show history with pagination
// @Tags history
// @Security UserID
// @Param page query int false "page number" minimum(1) default(1)
// @Param size query int false "page size" minimum(10) maximum(100) default(10)
// @Produce json
// @Success 200 {array} dto.HistoryResponse
// @Header 200 {int} X-Total-Count "500"
// @Router /history [get]
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

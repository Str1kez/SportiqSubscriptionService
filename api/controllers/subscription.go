package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type eventIdPath struct {
	Id string `uri:"event_id" binding:"required,uuid"`
}

func (ctl *Controller) Subscribe(c *gin.Context) {
	eventId := eventIdPath{}
	if err := c.ShouldBindUri(&eventId); err != nil {
		ctl.errorResponse(c, http.StatusUnprocessableEntity, "subscription.event.id.invalid", err)
		return
	}
	if err := ctl.subscriptionDB.Subscribe(c.GetHeader("User"), eventId.Id); err != nil {
		ctl.errorResponse(c, http.StatusBadRequest, "subscription.subscribe", err)
		return
	}
	ctl.response(c, http.StatusCreated, nil)
}

func (ctl *Controller) Unsubscribe(c *gin.Context) {
	eventId := eventIdPath{}
	if err := c.ShouldBindUri(&eventId); err != nil {
		ctl.errorResponse(c, http.StatusUnprocessableEntity, "subscription.event.id.invalid", err)
		return
	}
	if err := ctl.subscriptionDB.Unsubscribe(c.GetHeader("User"), eventId.Id); err != nil {
		ctl.errorResponse(c, http.StatusBadRequest, "subscription.unsubscribe", err)
		return
	}
	ctl.response(c, http.StatusOK, nil)
}

func (ctl *Controller) SubscribersCount(c *gin.Context) {
	eventId := eventIdPath{}
	if err := c.ShouldBindUri(&eventId); err != nil {
		ctl.errorResponse(c, http.StatusUnprocessableEntity, "subscription.event.id.invalid", err)
		return
	}
	var count uint
	count, err := ctl.subscriptionDB.CountSubscribers(eventId.Id)
	if err != nil {
		ctl.errorResponse(c, http.StatusBadRequest, "subscription.subscribers_count", err)
		return
	}
	ctl.response(c, http.StatusOK, gin.H{"event_id": eventId.Id, "subscribersCount": count})
}

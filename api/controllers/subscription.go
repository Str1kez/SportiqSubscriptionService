package controllers

import (
	"net/http"

	"github.com/Str1kez/SportiqSubscriptionService/api/responses"
	"github.com/gin-gonic/gin"
)

type eventIdPath struct {
	Id string `uri:"event_id" binding:"required,uuid"`
}

// @Summary Subscribe
// @ID subscribe
// @Description Subscription on event
// @Tags subscription
// @Security UserID
// @Param event_id path string true "UUID of event"
// @Produce json
// @Success 201 "subscribed"
// @Failure 422 {object} responses.ErrorResponse "invalid id of event"
// @Failure 400 {object} responses.ErrorResponse "subscription is unavailable"
// @Router /subscribe/{event_id} [post]
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

// @Summary Unsubscribe
// @ID unsubscribe
// @Description Unsubscription from event
// @Tags subscription
// @Security UserID
// @Param event_id path string true "UUID of event"
// @Produce json
// @Success 200 "success"
// @Failure 422 {object} responses.ErrorResponse "invalid id of event"
// @Failure 400 {object} responses.ErrorResponse "subscription is unavailable"
// @Router /unsubscribe/{event_id} [post]
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

// @Summary Subscriptions Count
// @ID subcount
// @Description Returns count of subscribers on event
// @Tags subscription
// @Security UserID
// @Param event_id path string true "UUID of event"
// @Produce json
// @Success 200 {object} responses.SubscriptionCountResponse "success"
// @Failure 422 {object} responses.ErrorResponse "invalid id of event"
// @Failure 400 {object} responses.ErrorResponse "count is unavailable"
// @Router /subscribers/count/{event_id} [get]
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
	ctl.response(c, http.StatusOK, responses.SubscriptionCountResponse{EventId: eventId.Id, SubscribersCount: count})
}

// @Summary Subscription Info
// @ID subinfo
// @Description Returns info about current state of subscriptions
// @Tags subscription
// @Security UserID
// @Produce json
// @Success 200 {array} dto.SubscriptionResponse "success"
// @Failure 400 {object} responses.ErrorResponse "subscriptions are unavailable"
// @Router /subscriptions [get]
func (ctl *Controller) Subscriptions(c *gin.Context) {
	subscriptions, err := ctl.subscriptionDB.GetEvents(c.GetHeader("User"))
	if err != nil {
		ctl.errorResponse(c, http.StatusBadRequest, "subscription.subscriptions", err)
		return
	}
	ctl.response(c, http.StatusOK, subscriptions)
}

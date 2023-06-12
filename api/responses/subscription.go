package responses

type SubscriptionCountResponse struct {
	EventId          string `json:"event_id"`
	SubscribersCount uint   `json:"subscribersCount"`
}

package responses

type SubscriptionCountResponse struct {
	EventId          string `json:"eventId"`
	SubscribersCount uint   `json:"subscribersCount"`
}

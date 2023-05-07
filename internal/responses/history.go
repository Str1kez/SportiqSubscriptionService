package responses

type HistoryResponse struct {
	EventId    string `json:"event_id"`
	EventTitle string `json:"event_title"`
	UserId     string `json:"user_id"`
	IsDeleted  bool   `json:"is_deleted"`
}

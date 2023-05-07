package responses

type HistoryResponse struct {
	EventId    string `json:"event_id" db:"id"`
	EventTitle string `json:"event_title" db:"title"`
	UserId     string `json:"user_id,omitempty"`
	IsDeleted  bool   `json:"is_deleted" db:"is_deleted"`
}

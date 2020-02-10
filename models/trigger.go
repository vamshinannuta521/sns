package models

type Trigger struct {
	ID        int    `json:"id"`
	EventID   string `json:"event_id"`
	AccountID string `json:"account_id"`
	Message   string `json:"message"`
}

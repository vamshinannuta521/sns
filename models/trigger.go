package models

type Trigger struct {
	ID        int    `json:"id"`
	EventID   int    `json:"event_id"`
	AccountID int    `json:"account_id"`
	Message   string `json:"message"`
}

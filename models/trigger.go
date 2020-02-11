package models

type Trigger struct {
	ID          string `json:"id"`
	EventName   string `json:"event_name"`
	AccountName string `json:"created_by"`
	Message     string `json:"message"`
}

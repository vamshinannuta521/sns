package models

type Action struct {
	ID          int    `json:"id"`
	EventName   string `json:"event_name"`
	ActionType  string `json:"action_type"`
	ActionSpec  string `json:"action_spec"`
	AccountName string `json:"created_by"`
}

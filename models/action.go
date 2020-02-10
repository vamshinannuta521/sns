package models

type Action struct {
	ID         int    `json:"id"`
	EventID    int    `json:"event_id"`
	ActionType string `json:"action_type"`
	ActionSpec string `json:"action_spec"`
	AccountID  int    `json:"account_id"`
}

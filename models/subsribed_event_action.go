package models

type SubscribedEventAction struct {
	ID         int    `json:"id"`
	EventID    string `json:"event_id"`
	ActionType int    `json:"action_type"`
	ActionSpec string `json:"action_spec"`
	AccountID  string `json:"account_id:`
}

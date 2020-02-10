package models

type Event struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AccountID string `json:"created_by"`
}

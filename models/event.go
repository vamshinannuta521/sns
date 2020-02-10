package models

type Event struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AccountID int    `json:"created_by"`
}

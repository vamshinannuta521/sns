package models

type Event struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	AccountName string `json:"created_by"`
}

package models

type Event struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedBy string `json:"created_by"`
}

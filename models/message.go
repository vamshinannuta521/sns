package models

import (
	"time"
)

type Trigger struct {
	ID        int       `json:"id"`
	EventID   string    `json:"event_id"`
	CreatedBy string    `json:"created_by"`
	Message   string    `json:"message"`
	CreatedOn time.Time `json:"created_on"`
}

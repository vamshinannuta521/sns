package event

import (
	"encoding/json"
	"fmt"
	da "sns/dataaccess"
	model "sns/models"
)

type SvcInterface interface {
	Get() string
	RegisterEvent([]byte) string
	GetEventsList() string
}

type Svc struct {
}

func NewSvc() *Svc {
	return &Svc{}
}

func (s *Svc) RegisterEvent(req []byte) string {
	var event model.Event
	json.Unmarshal(req, &event)
	fmt.Println("register event", event)
	da.CreateEvent(&event)
	return "event registered"

}

func (s *Svc) Get() string {
	return "single event"
}

func (s *Svc) GetEventsList() string {
	return "events list"
}

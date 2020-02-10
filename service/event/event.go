package event

import (
	"fmt"
	"sns/dataaccess"
	model "sns/models"

	"github.com/sirupsen/logrus"
)

type SvcInterface interface {
	Get(string) string
	RegisterEvent(model.Event) string
	GetEventsList() string
}

var logger *logrus.Entry

type Event struct {
	Name      string `json:"name"`
	CreatedBy string `json:"created_by"`
}

type Svc struct {
	*dataaccess.PostgresClient
}

func NewSvc(client *dataaccess.PostgresClient, log *logrus.Entry) *Svc {
	logger = log
	return &Svc{
		PostgresClient: client,
	}
}

func (s *Svc) RegisterEvent(event model.Event) string {
	//var event model.Event
	//json.Unmarshal(req, &event)
	logger.Infof("Registering event", event)
	s.CreateEvent(&event)
	return "event registered"

}

func (s *Svc) Get(id string) string {
	fmt.Println("Returning Event of Id ", id)
	return "single event"
}

func (s *Svc) GetEventsList() string {
	return "events list"
}

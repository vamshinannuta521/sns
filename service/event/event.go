package event

import (
<<<<<<< HEAD
	"encoding/json"
	"fmt"
	da "sns/dataaccess"
	model "sns/models"
)

type SvcInterface interface {
	Get() string
	RegisterEvent([]byte) string
	GetEventsList() string
=======
	"sns/dataaccess"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Entry

type SvcInterface interface {
	// Get() string
	// Create([]byte)
>>>>>>> master
}

type Event struct {
	Name      string `json:"name"`
	CreatedBy string `json:"created_by"`
}

type Svc struct {
	*dataaccess.PostgresClient
}

<<<<<<< HEAD
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
=======
func NewSvc(client *dataaccess.PostgresClient, log *logrus.Entry) *Svc {
	logger = log
	return &Svc{
		PostgresClient: client,
	}
>>>>>>> master
}

// func (s *Svc) Create(c []byte) string {
// 	var e Event
// 	json.Unmarshal(c, e)
// 	e
// 	return "event"
// }

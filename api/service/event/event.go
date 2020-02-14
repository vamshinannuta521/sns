package event

import (
	"sns/dataaccess"
	"sns/models"
	model "sns/models"

	"github.com/sirupsen/logrus"
)

type SvcInterface interface {
	Get(string) (*models.Event, error)
	RegisterEvent(model.Event) error
	GetEventsList() ([]*models.Event, error)
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

func (s *Svc) RegisterEvent(event model.Event) error {
	logger.Infof("Registering event", event)
	return s.CreateEvent(&event)

}

func (s *Svc) Get(id string) (*models.Event, error) {
	logger.Infof("Returning Event of ID", id)
	event, err := s.GetEvent(id)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return event, nil
}

func (s *Svc) GetEventsList() ([]*models.Event, error) {
	logger.Info("Returning list of events")
	events, err := s.GetAllEvents()
	if err != nil {
		logger.Error(err)

		return nil, err
	}
	return events, nil

}

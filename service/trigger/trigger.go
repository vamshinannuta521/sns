package trigger

import (
	"sns/dataaccess"
	"sns/models"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Entry

type SvcInterface interface {
	Get(string) (*models.Trigger, error)
	Create(*models.Trigger) error
	GetList() ([]*models.Trigger, error)
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

func (s *Svc) Create(trigger *models.Trigger) error {
	err := s.CreateTrigger(trigger)
	logger.Error(err)
	return err

}

func (s *Svc) Get(triggerID string) (*models.Trigger, error) {
	trigger, err := s.GetTrigger(triggerID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return trigger, nil

}

func (s *Svc) GetList() ([]*models.Trigger, error) {
	triggers, err := s.GetAllTriggers()
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return triggers, nil

}

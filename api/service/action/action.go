package action

import (
	"sns/dataaccess"
	"sns/models"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Entry

type SvcInterface interface {
	Get(string) (*models.Action, error)
	Create(*models.Action) error
	GetList() ([]*models.Action, error)
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

func (s *Svc) Create(action *models.Action) error {
	err := s.CreateAction(action)
	logger.Error(err)
	return err

}

func (s *Svc) Get(actionID string) (*models.Action, error) {
	action, err := s.GetAction(actionID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return action, nil

}

func (s *Svc) GetList() ([]*models.Action, error) {
	actions, err := s.GetAllActions()
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return actions, nil

}

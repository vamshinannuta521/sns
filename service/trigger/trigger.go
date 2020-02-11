package trigger

import (
	"sns/dataaccess"
	"sns/models"

	"github.com/satori/go.uuid"
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

	ch chan *models.Trigger
}

func NewSvc(client *dataaccess.PostgresClient, log *logrus.Entry) *Svc {
	logger = log
	svc := &Svc{
		PostgresClient: client,
		ch:             make(chan *models.Trigger, 100),
	}

	workers := 2
	for i := 0; i < workers; i++ {
		go svc.PushToKafka()
	}

	return svc
}

func (s *Svc) Create(trigger *models.Trigger) error {
	trigger.ID = NewUUID()
	err := s.CreateTrigger(trigger)
	if err != nil {
		logger.Error(err)
		return err
	}
	s.ch <- trigger
	return nil

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

func (s *Svc) PushToKafka() {
	for trigger := range s.ch {
		actions, err := s.GetAllActionsWithEventFilter(trigger.EventName)
		if err != nil {
			logger.Error(err)
			continue
		}
		logger.Info(actions)

	}
}

//NewUUID return new uuid
func NewUUID() string {
	UUID := uuid.NewV4()
	return UUID.String()
}

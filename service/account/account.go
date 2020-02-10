package account

import (
	"sns/dataaccess"
	"sns/models"

	"github.com/sirupsen/logrus"
)

var log = logrus.NewEntry(logrus.New())

type SvcInterface interface {
	Get(string) (*models.Account, error)
	Create(*models.Account) error
}

type Svc struct {
	*dataaccess.PostgresClient
}

func NewSvc() *Svc {
	client, _ := dataaccess.NewClient()

	return &Svc{
		PostgresClient: client,
	}
}

func (s *Svc) Create(account *models.Account) error {
	err := s.CreateAccount(account)
	log.Error(err)
	return err

}

func (s *Svc) Get(accountID string) (*models.Account, error) {
	account, err := s.GetAccount(accountID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return account, nil

}

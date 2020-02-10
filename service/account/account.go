package account

import (
	"sns/dataaccess"
	"sns/models"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Entry

type SvcInterface interface {
	Get(string) (*models.Account, error)
	Create(*models.Account) error
	GetList() ([]*models.Account, error)
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

func (s *Svc) Create(account *models.Account) error {
	err := s.CreateAccount(account)
	logger.Error(err)
	return err

}

func (s *Svc) Get(accountID string) (*models.Account, error) {
	account, err := s.GetAccount(accountID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return account, nil

}

func (s *Svc) GetList() ([]*models.Account, error) {
	accounts, err := s.GetAllAccounts()
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return accounts, nil

}

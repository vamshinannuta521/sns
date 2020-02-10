package account

import (
	"sns/dataaccess"
	"sns/models"

	"github.com/sirupsen/logrus"
)

var log = logrus.NewEntry(logrus.New())

type SvcInterface interface {
	Get(string) (*models.Account, error)
	Create([]byte) error
}

type Svc struct {
}

func NewSvc() *Svc {
	return &Svc{}
}

func (s *Svc) Create(payload []byte) error {
	var account models.Account
	err := json.Unmarshal(payload, &account)
	if err != nil {
		log.Error(err)
		return err
	}
	err = dataaccess.CreateAccount(account)
	return err

}

func (s *Svc) Get(accountID string) (*models.Account, error) {
	account, err := dataaccess.GetAccount(accountID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return account, nil

}

package account

import (
	"sns/dataaccess"

	"github.com/sirupsen/logrus"
)

var log = logrus.NewEntry(logrus.New())

type SvcInterface interface {
	Get(string) Account
	Create([]byte) error
}

type Svc struct {
}

type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewSvc() *Svc {
	return &Svc{}
}

func (s *Svc) Create(payload []byte) error {
	var account Account
	err := json.Unmarshal(payload, &account)
	if err != nil {
		log.Error(err)
		return err
	}

}

func (s *Svc) Get(accountID string) error {
	var account Account
	err := json.Unmarshal(payload, &account)
	if err != nil {
		log.Error(err)
		return err
	}

}

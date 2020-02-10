package action

import (
	"sns/dataaccess"
	// "sns/models"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Entry

type SvcInterface interface {
	// Get(string) (*models.SubscribedEventAction, error)
	// Create([]byte) error
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

// func (s *Svc) Create(payload []byte) error {
// 	var account models.Account
// 	err := json.Unmarshal(payload, &account)
// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}
// 	err = dataaccess.CreateAccount(account)
// 	return err

// }

// func (s *Svc) Get(accountID string) (*models.Account, error) {
// 	account, err := dataaccess.GetAccount(accountID)
// 	if err != nil {
// 		log.Error(err)
// 		return nil, err
// 	}
// 	return account, nil

// }

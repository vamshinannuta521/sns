package trigger

import (
	"sns/dataaccess"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Entry

type SvcInterface interface {
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

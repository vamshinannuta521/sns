package event

import (
	"sns/dataaccess"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Entry

type SvcInterface interface {
	// Get() string
	// Create([]byte)
}

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

// func (s *Svc) Create(c []byte) string {
// 	var e Event
// 	json.Unmarshal(c, e)
// 	e
// 	return "event"
// }

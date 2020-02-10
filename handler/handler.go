package handler

import (
	"fmt"
	"net/http"

	"sns/models"
	"sns/service/account"
	"sns/service/action"
	"sns/service/event"
	"sns/service/trigger"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	eventSvc   event.SvcInterface
	accountSvc account.SvcInterface
	actionSvc  action.SvcInterface
	triggerSvc trigger.SvcInterface
}

var logger *logrus.Entry

func NewHandler(eventSvc event.SvcInterface,
	accountSvc account.SvcInterface,
	actionSvc action.SvcInterface,
	triggerSvc trigger.SvcInterface,
	log *logrus.Entry) *Handler {

	logger = log

	return &Handler{
		eventSvc:   eventSvc,
		accountSvc: accountSvc,
		actionSvc:  actionSvc,
		triggerSvc: triggerSvc,
	}
}

//DefaultHandler handler
func (handler *Handler) Default(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "success done")
}

//DefaultHandler handler
func (handler *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {

	err := handler.accountSvc.Create(&models.Account{
		Name: "vamshi",
	})
	logger.Error(err)
	fmt.Fprint(w, "success done")
}

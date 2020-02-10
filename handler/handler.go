package handler

import (
	"fmt"
	"net/http"

	"sns/models"
	"sns/service/account"
	"sns/service/action"
	"sns/service/event"
	"sns/service/trigger"
)

type Handler struct {
	eventSvc   event.SvcInterface
	accountSvc account.SvcInterface
	actionSvc  action.SvcInterface
	triggerSvc trigger.SvcInterface
}

func NewHandler(eventSvc event.SvcInterface,
	accountSvc account.SvcInterface,
	actionSvc action.SvcInterface,
	triggerSvc trigger.SvcInterface) *Handler {
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
	fmt.Fprint(w, "success done")
}

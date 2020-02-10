package handler

import (
	"fmt"
	"net/http"

	"sns/service/account"
	"sns/service/action"
	"sns/service/event"
	"sns/service/trigger"
)

type Handler struct {
	eventSvc   event.EventSvcInterface
	accountSvc account.AccountSvcInterface
	actionSvc  action.ActionSvcInterface
	triggerSvc trigger.TriggerSvcInterface
}

func NewHandler(eventSvc event.EventSvcInterface,
	accountSvc account.AccountSvcInterface,
	actionSvc action.ActionSvcInterface,
	triggerSvc trigger.TriggerSvcInterface) *Handler {
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
func (handler *Handler) GetEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, handler.eventSvc.Get())
}

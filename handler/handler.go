package handler

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"net/http"
	"sns/models"
	model "sns/models"
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

func CheckContentType(r *http.Request) bool {
	if r.Header.Get("Content-Type") != "" {
		if r.Header.Get("Content-Type") != "application/json" {
			return false
		}
	} else {
		return false
	}
	return true

}

//Default handler
func (handler *Handler) Default(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "success done")
	uuid := returnUUID(r)
	fmt.Println(uuid)
}

//GetEvent handler
func (handler *Handler) GetEvent(w http.ResponseWriter, r *http.Request) {
	eventId := returnUUID(r)
	fmt.Fprint(w, handler.eventSvc.Get(eventId))
}

//DefaultHandler handler
func (handler *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {

	err := handler.accountSvc.Create(&models.Account{
		Name: "vamshi",
	})
	logger.Error(err)
	fmt.Fprint(w, "success done")
}

//GetEventsList handler
func (handler *Handler) GetEventsList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, handler.eventSvc.GetEventsList())
}

//RegisterEventHandler
func (handler *Handler) RegisterEvent(w http.ResponseWriter, r *http.Request) {
	var event model.Event
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&event)
	if CheckContentType(r) {
		fmt.Fprint(w, handler.eventSvc.RegisterEvent(event))

	} else {
		fmt.Println("Not application/json")
	}

}

//TriggerEventHandler
func (handler *Handler) TriggerEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("triggering event ..")
}

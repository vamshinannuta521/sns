package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sns/models"
	model "sns/models"
	"sns/service/account"
	"sns/service/action"
	"sns/service/event"
	"sns/service/trigger"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	successResponse = `success`
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
	sendSuccessResponse(w)
}

func (handler *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {

	var account models.Account
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&account)

	err := handler.accountSvc.Create(&account)
	if err != nil {
		logger.Error(err)
		sendInternalServerErrorResponse(w, err)
		return
	}

	sendSuccessResponse(w)

}

func (handler *Handler) GetAccount(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	accountID, ok := pathParams["uuid"]
	if !ok {
		sendBadRequestResponse(w)
		return
	}
	account, err := handler.accountSvc.Get(accountID)
	if err != nil {
		sendInternalServerErrorResponse(w, err)
		return
	}

	resp, err := json.Marshal(account)
	if err != nil {
		sendInternalServerErrorResponse(w, err)
		return
	}
	w.Write(resp)
}

func (handler *Handler) GetAccountList(w http.ResponseWriter, r *http.Request) {
	accounts, err := handler.accountSvc.GetList()
	if err != nil {
		sendInternalServerErrorResponse(w, err)
		return
	}

	resp, err := json.Marshal(accounts)
	if err != nil {
		sendInternalServerErrorResponse(w, err)
		return
	}
	logger.Infof("Returning accounts list")
	w.Write(resp)
}

//GetEvent handler
func (handler *Handler) GetEvent(w http.ResponseWriter, r *http.Request) {
	eventId := returnUUID(r)
	event, err := handler.eventSvc.Get(eventId)
	if err != nil {
		sendInternalServerErrorResponse(w, err)
		return
	}
	resp, err := json.Marshal(event)
	if err != nil {
		sendInternalServerErrorResponse(w, err)
		return
	}
	w.Write(resp)
}

//GetEventsList handler
func (handler *Handler) GetEventsList(w http.ResponseWriter, r *http.Request) {
	events, err := handler.eventSvc.GetEventsList()
	if err != nil {
		sendInternalServerErrorResponse(w, err)
		return
	}
	resp, err := json.Marshal(events)
	if err != nil {
		sendInternalServerErrorResponse(w, err)
		return
	}
	logger.Infof("Returning events list")
	w.Write(resp)
}

//RegisterEventHandler
func (handler *Handler) RegisterEvent(w http.ResponseWriter, r *http.Request) {
	var event model.Event
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&event)
	if CheckContentType(r) {
		err := handler.eventSvc.RegisterEvent(event)
		if err != nil {
			logger.Error(err)
			sendInternalServerErrorResponse(w, err)
			return
		}

	} else {
		sendInternalServerErrorResponse(w, fmt.Errorf("Content type not application/json"))
		return
	}

	sendSuccessResponse(w)

}

//TriggerEventHandler
func (handler *Handler) TriggerEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("triggering event ..")
}

func (handler *Handler) CreateAction(w http.ResponseWriter, r *http.Request) {

	var action models.Action
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&action)

	err := handler.actionSvc.Create(&action)
	if err != nil {
		logger.Error(err)
		sendInternalServerErrorResponse(w, err)
		return
	}

	sendSuccessResponse(w)

}

func (handler *Handler) GetAction(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	actionID, ok := pathParams["uuid"]
	if !ok {
		sendBadRequestResponse(w)
		return
	}
	action, err := handler.actionSvc.Get(actionID)
	if err != nil {
		sendInternalServerErrorResponse(w, err)
		return
	}

	resp, err := json.Marshal(action)
	if err != nil {
		sendInternalServerErrorResponse(w, err)
		return
	}
	w.Write(resp)
}

func (handler *Handler) GetActionList(w http.ResponseWriter, r *http.Request) {

	actions, err := handler.actionSvc.GetList()
	if err != nil {
		sendInternalServerErrorResponse(w, err)
		return
	}

	resp, err := json.Marshal(actions)
	if err != nil {
		sendInternalServerErrorResponse(w, err)
		return
	}
	w.Write(resp)
}

func (handler *Handler) CreateTrigger(w http.ResponseWriter, r *http.Request) {

	var trigger models.Trigger
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&trigger)

	err := handler.triggerSvc.Create(&trigger)
	if err != nil {
		logger.Error(err)
		sendInternalServerErrorResponse(w, err)
		return
	}

	sendSuccessResponse(w)

}

func (handler *Handler) GetTrigger(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	triggerID, ok := pathParams["uuid"]
	if !ok {
		sendBadRequestResponse(w)
		return
	}
	trigger, err := handler.triggerSvc.Get(triggerID)
	if err != nil {
		sendInternalServerErrorResponse(w, err)
		return
	}

	resp, err := json.Marshal(trigger)
	if err != nil {
		sendInternalServerErrorResponse(w, err)
		return
	}
	w.Write(resp)
}

func (handler *Handler) GetTriggerList(w http.ResponseWriter, r *http.Request) {

	triggers, err := handler.triggerSvc.GetList()
	if err != nil {
		sendInternalServerErrorResponse(w, err)
		return
	}

	resp, err := json.Marshal(triggers)
	if err != nil {
		sendInternalServerErrorResponse(w, err)
		return
	}
	w.Write(resp)
}

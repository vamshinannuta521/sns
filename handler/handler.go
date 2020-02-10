package handler

import (
	"encoding/json"
	"net/http"

	"sns/models"
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

//DefaultHandler handler
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
	w.Write(resp)
}

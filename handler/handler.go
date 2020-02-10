package handler

import (
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

// func CheckContentType(r *http.Request) {
// 	if r.Header.Get("Content-Type") != "" {
// 		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
// 		if value != "application/json" {
// 			return false
// 		}
// 	} else {
// 		return false
// 	}
// 	return true

// }

//Default handler
func (handler *Handler) Default(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "success done")
	uuid := returnUUID(r)
	fmt.Println(uuid)
}

//GetEvent handler
func (handler *Handler) GetEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, handler.eventSvc.Get())
}

//GetEventsList handler
func (handler *Handler) GetEventsList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, handler.eventSvc.GetEventsList())
}

//RegisterEventHandler
func (handler *Handler) RegisterEvent(w http.ResponseWriter, r *http.Request) {
	rBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprint(w, handler.eventSvc.RegisterEvent(rBody))
	//if CheckContentType(r) {
	// event := models.CreateEvent(r)

	// } else {
	// 	fmt.Println("Not application/json")
	// }
	//da.CreateEvent(&event)

}

//TriggerEventHandler
func (handler *Handler) TriggerEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("triggering event ..")
}

package handler

import (
	"fmt"
	"net/http"

	"sns/service/event"
)

type Handler struct {
	eventSvc event.EventSvc
}

func NewHandler(eSvc event.EventSvc) *Handler {
	return &Handler{
		eventSvc: eSvc,
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

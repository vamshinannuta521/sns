package handler

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

//DefaultHandler handler
func (handler *Handler) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "success done")
}

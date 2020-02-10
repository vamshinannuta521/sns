package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

//DefaultHandler handler for plugin mandatory routes other than execute
func (handler *Handler) DefaultHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "success done")
}

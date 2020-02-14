package handler

import (
	"fmt"
	"net/http"
	"path"
)

func sendSuccessResponse(w http.ResponseWriter) {
	msg := `{"message":"success"}`
	w.Write([]byte(msg))
	return
}

func sendInternalServerErrorResponse(w http.ResponseWriter, err error) {
	msg := fmt.Sprintf(`{"message":"%s"}`, err.Error())
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(msg))
	return
}

func sendBadRequestResponse(w http.ResponseWriter) {
	msg := fmt.Sprintf(`{"message":"%s"}`, "Bad request")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(msg))
	return

}

func returnUUID(r *http.Request) string {
	return path.Base(r.URL.Path)
}

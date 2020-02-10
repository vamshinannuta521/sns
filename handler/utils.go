package handler

import (
	"net/http"
	"path"
)

func returnUUID(r *http.Request) string {
	return path.Base(r.URL.Path)
}

package routes

import (
	"net/http"
	"sns/handler"

	"github.com/gorilla/mux"
)

type Router struct {
	router  *mux.Router
	handler *handler.Handler
}

//NewRouter return new router
func NewRouter(handler *handler.Handler) *Router {
	router := &Router{
		router:  mux.NewRouter(),
		handler: handler,
	}
	return router
}

//GetRouter return router
func (r *Router) GetRouter() *mux.Router {
	return r.router
}

//ConfigureRoutes plugin routes
func (r *Router) ConfigureRoutes() {

	r.router.HandleFunc("/sns/api/v1/account", r.handler.CreateAccount).Methods(http.MethodPost)
	r.router.HandleFunc("/sns/api/v1/account/{uuid}", r.handler.Default).Methods(http.MethodGet)
	r.router.HandleFunc("/sns/api/v1/event", r.handler.Default).Methods(http.MethodPost)
	r.router.HandleFunc("/sns/api/v1/event/{uuid}", r.handler.Default).Methods(http.MethodGet)
	r.router.HandleFunc("/sns/api/v1/event/list", r.handler.Default).Methods(http.MethodGet)
	r.router.HandleFunc("/sns/api/v1/subscribe", r.handler.Default).Methods(http.MethodPost)
	r.router.HandleFunc("/sns/api/v1/trigger", r.handler.TriggerEvent).Methods(http.MethodPost)

}

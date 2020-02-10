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
	r.router.HandleFunc("/sns/api/v1/account/list", r.handler.GetAccountList).Methods(http.MethodGet)
	r.router.HandleFunc("/sns/api/v1/account/{uuid}", r.handler.GetAccount).Methods(http.MethodGet)
	r.router.HandleFunc("/sns/api/v1/event", r.handler.RegisterEvent).Methods(http.MethodPost)
	r.router.HandleFunc("/sns/api/v1/event/list", r.handler.GetEventsList).Methods(http.MethodGet)
	r.router.HandleFunc("/sns/api/v1/event/{uuid}", r.handler.GetEvent).Methods(http.MethodGet)
	r.router.HandleFunc("/sns/api/v1/action", r.handler.CreateAction).Methods(http.MethodPost)
	r.router.HandleFunc("/sns/api/v1/action/list", r.handler.GetActionList).Methods(http.MethodGet)
	r.router.HandleFunc("/sns/api/v1/action/{uuid}", r.handler.GetAction).Methods(http.MethodGet)
	r.router.HandleFunc("/sns/api/v1/trigger", r.handler.CreateTrigger).Methods(http.MethodPost)
	r.router.HandleFunc("/sns/api/v1/trigger/list", r.handler.GetTriggerList).Methods(http.MethodGet)
	r.router.HandleFunc("/sns/api/v1/trigger/{uuid}", r.handler.GetTrigger).Methods(http.MethodGet)
	r.router.HandleFunc("/sns/api/v1/trigger/{event}", r.handler.TriggerEvent).Methods(http.MethodPost)
	r.router.HandleFunc("", r.handler.Default)

}

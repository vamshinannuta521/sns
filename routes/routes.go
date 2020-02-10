package routes

import (
	"sns/handler"

	"github.com/julienschmidt/httprouter"
)

type Router struct {
	router  *httprouter.Router
	handler *handler.Handler
}

//NewRouter return new router
func NewRouter(handler *handler.Handler) *Router {
	router := &Router{
		router:  httprouter.New(),
		handler: handler,
	}
	return router
}

//GetRouter return router
func (r *Router) GetRouter() *httprouter.Router {
	return r.router
}

//ConfigureRoutes plugin routes
func (r *Router) ConfigureRoutes() {

	r.router.POST("/sns/api/v1/user", r.handler.DefaultHandler)
	r.router.GET("/sns/api/v1/user/:uuid", r.handler.DefaultHandler)
	r.router.POST("/sns/api/v1/event", r.handler.DefaultHandler)
	r.router.GET("/sns/api/v1/event/:uuid", r.handler.DefaultHandler)
	r.router.POST("/sns/api/v1/event/list", r.handler.DefaultHandler)
	r.router.POST("/sns/api/v1/execute/:req_id", r.handler.DefaultHandler)
	r.router.POST("/sns/api/v1/subscribe", r.handler.DefaultHandler)
	r.router.POST("/sns/api/v1/trigger/:event", r.handler.DefaultHandler)

}

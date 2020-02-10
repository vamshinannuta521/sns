package main

import (
	"fmt"
	"net/http"

	"sns/handler"
	"sns/routes"
	"sns/service/account"
	"sns/service/action"
	"sns/service/event"
	"sns/service/trigger"
)

func main() {
	fmt.Println("hi")

	eventSvc := event.NewEventSvc()
	accountSvc := event.NewEventSvc()
	actionSvc := event.NewEventSvc()
	triggerSvc := event.NewEventSvc()

	handler := handler.NewHandler(eventSvc, accountSvc, actionSvc, triggerSvc)

	//initialize router with handler
	router := routes.NewRouter(handler)
	router.ConfigureRoutes()

	//start server
	http.ListenAndServe(fmt.Sprintf("%s:%d", "localhost", 6666), router.GetRouter())

}

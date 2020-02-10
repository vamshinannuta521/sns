package main

import (
	"fmt"
	"net/http"

	"sns/handler"
	"sns/routes"
	"sns/service/event"
)

func main() {
	fmt.Println("hi")

	eventSvc := event.NewEventSvc()

	handler := handler.NewHandler(eventSvc)

	//initialize router with handler
	router := routes.NewRouter(handler)
	router.ConfigureRoutes()

	//start server
	http.ListenAndServe(fmt.Sprintf("%s:%d", "localhost", 6666), router.GetRouter())

}

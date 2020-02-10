package main

import (
	"fmt"
	"net/http"

	"sns/handler"
	"sns/routes"
)

func main() {
	fmt.Println("hi")

	handler := handler.NewHandler()

	//initialize router with handler
	router := routes.NewRouter(handler)
	router.ConfigureRoutes()

	//start server
	http.ListenAndServe(fmt.Sprintf("%s:%d", "localhost", 6666), router.GetRouter())

}

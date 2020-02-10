package main

import (
	"fmt"
	"net/http"

	"sns/dataaccess"
	"sns/handler"
	"sns/routes"
	"sns/service/account"
	"sns/service/action"
	"sns/service/event"
	"sns/service/trigger"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.NewEntry(logrus.New())
}

func main() {
	fmt.Println("main started")

	dbclient, err := dataaccess.NewClient(log)
	if err != nil {
		log.Fatal(err)
	}

	eventSvc := event.NewSvc(dbclient, log)
	accountSvc := account.NewSvc(dbclient, log)
	actionSvc := action.NewSvc(dbclient, log)
	triggerSvc := trigger.NewSvc(dbclient, log)

	handler := handler.NewHandler(eventSvc, accountSvc, actionSvc, triggerSvc, log)

	//initialize router with handler
	router := routes.NewRouter(handler)
	router.ConfigureRoutes()

	//start server
	http.ListenAndServe(fmt.Sprintf("%s:%d", "localhost", 6666), router.GetRouter())

}

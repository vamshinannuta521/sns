package main

import (
	"fmt"
	"httpnotificationprocessor/handler"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler.UpgradeHandler)
	err := http.ListenAndServe(":8888", nil)
	fmt.Println(err)
}

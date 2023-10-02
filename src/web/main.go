package main

import (
	"fmt"
	"net/http"
	"web/models"
)

func main() {
	handleRequest()
}

func handleRequest() {
	http.HandleFunc("/", getHomePage)
	http.HandleFunc("/about/", aboutPage)
	err := http.ListenAndServe("localhost:3543", nil)
	if err != nil {
		return
	}
}

func aboutPage(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}
	user.SetName("Adel")
	user.SetAge(19)
	user.SetMoney(5)
	_, err := fmt.Fprintf(writer, user.GetAllInfo())
	if err != nil {
		return
	}
}

func getHomePage(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "Go is super easy")
	if err != nil {
		return
	}
}

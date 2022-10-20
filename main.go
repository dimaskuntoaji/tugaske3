package main

import (
	"tugaske3/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/", controller.GetStatus)

	http.ListenAndServe(":8080", nil)
}
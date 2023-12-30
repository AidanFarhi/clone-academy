package app

import (
	"clone-academy/app/handlers"
	"net/http"
)

func setupHandlers() {
	http.HandleFunc("/", handlers.IndexHandler)
}

func Run() {
	setupHandlers()
	http.ListenAndServe(":8080", nil)
}

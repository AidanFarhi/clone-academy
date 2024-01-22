package app

import (
	"clone-academy/app/config"
	"clone-academy/app/handlers"
	"net/http"
)

func setupHandlers() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/courses", handlers.CoursesHandler)
	http.HandleFunc("/add-course", handlers.AddCourseHandler)
	http.HandleFunc("/add-course-page", handlers.AddCoursePageHandler)
	http.Handle("/static/", handlers.StaticHandler())
}

func Run() {
	config.InitDatabase()
	setupHandlers()
	http.ListenAndServe(":8080", nil)
}

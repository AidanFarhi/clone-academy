package handlers

import (
	"clone-academy/app/service"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "app/web/index.html")
}

func CoursesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("app/web/templates/courses.html")
	courses := service.GetCourses()
	tmpl.Execute(w, courses)
}

func StaticHandler() http.Handler {
	staticBaseDir := "/static/"
	fullPath := "app/web/static"
	handler := http.StripPrefix(staticBaseDir, http.FileServer(http.Dir(fullPath)))
	return handler
}

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

func AddCourseHandler(w http.ResponseWriter, r *http.Request) {
	service.AddCourse(&r.Form)
	IndexHandler(w, r)
}

func AddCoursePageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "app/web/templates/add_course.html")
}

func StaticHandler() http.Handler {
	staticBaseDir := "/static/"
	fullPath := "app/web/static"
	handler := http.StripPrefix(staticBaseDir, http.FileServer(http.Dir(fullPath)))
	return handler
}

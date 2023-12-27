package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	ID    int
	Title string
}

var todos []Todo

func main() {
	// Serve static files (optional - for demonstration purposes)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Define routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addTodoHandler)

	// Start the server
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML template and render it
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, todos)
}

func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	r.ParseForm()
	title := r.FormValue("title")

	// Create a new todo
	newTodo := Todo{
		ID:    len(todos) + 1,
		Title: title,
	}

	// Add the new todo to the list
	todos = append(todos, newTodo)

	// Redirect to the home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

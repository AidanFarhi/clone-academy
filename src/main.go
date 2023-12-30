package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
var pagesDirectory = "pages"
var templateDirectory = "templates"
var allTemplates *template.Template

func loadTemplates(templateDirectory string) (*template.Template, error) {
	entries, err := os.ReadDir(templateDirectory)
	if err != nil {
		return nil, err
	}
	templateFileNames := make([]string, 0)
	for _, entry := range entries {
		templateFileNames = append(templateFileNames, templateDirectory+"/"+entry.Name())
	}
	templates, err := template.ParseFiles(templateFileNames...)
	if err != nil {
		return nil, err
	}
	return templates, err
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := allTemplates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := pagesDirectory + "/" + p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := pagesDirectory + "/" + title + ".txt"
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: bytes}, nil
}

func main() {
	templates, err := loadTemplates(templateDirectory)
	if err != nil {
		log.Default().Panic(err.Error())
	}
	allTemplates = templates
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// web-server test
package main

import (
	"html/template"
	"net/http"
	"regexp"

	"github.com/AndreaGhizzoni/go-introduction/10-WebServer/page"
)

// Variable to hold all template file just specify the patter where to find all
// the templates
// Templates are parser when program is started, not at user request
var templates = template.Must(template.ParseGlob("template/*.html"))

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// Auxiliary function to render a specific template.
// w writer where to write the template
// tmpl template name
// p page structure where to find the data
func renderTemplate(w http.ResponseWriter, tmpl string, p *page.Structure) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Http handler for index
func rootHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", nil)
}

// Http handler for view template
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := page.Load(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

// Http handler for edit template
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := page.Load(title)
	if err != nil {
		p = &page.Structure{Title: title}
	}
	renderTemplate(w, "edit", p)
}

// Http handler for save template
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &page.Structure{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// Function made to avoid the validation path for each handler (except root)
// This function take a function fn(http.ResponseWriter, *http.Request, string)
// as argument, and return a function f( w http.ResponseWriter, r *http.Request)
// that validate the path r and call fn(w, r, m[2])
// Note that the return function has the same signature as the second parameter
// of http.HandleFunc() used in main()
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

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.ListenAndServe(":8080", nil)
}

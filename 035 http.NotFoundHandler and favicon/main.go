package main

import (
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/favicon.ico")
}

func init() {
	tpl = template.Must(template.ParseFiles("templates/tpl.gohtml"))
}

var tpl *template.Template

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/favicon.ico", favicon)

	http.Handle("/random/url", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

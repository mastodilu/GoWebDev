package main

import (
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	heSaid := r.FormValue("text")
	tpl.Execute(w, heSaid)
}

func init() {
	tpl = template.Must(template.ParseFiles("templates/tpl.gohtml"))
}

var tpl *template.Template

func main() {
	http.HandleFunc("/", home)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

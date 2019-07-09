package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request) {
	tpl.Execute(res, nil)
}

func init() {
	tpl = template.Must(template.ParseFiles("templates/tpl.gohtml"))
}

var tpl *template.Template

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./public"))))
	http.HandleFunc("/", home)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

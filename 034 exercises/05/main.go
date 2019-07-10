package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func apply(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	if req.Method == http.MethodPost {
		applyProcess(res, req)
		return
	}

	err := tpl.ExecuteTemplate(res, "apply.gohtml", nil)
	HandleError(res, err)
}

func applyProcess(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
	HandleError(w, err)
}

// HandleError -
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

var tpl *template.Template

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about/", about)
	http.HandleFunc("/apply/", apply)
	http.HandleFunc("/contact/", contact)

	http.ListenAndServe(":8080", nil)
}

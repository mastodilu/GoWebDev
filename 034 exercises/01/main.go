package main

import (
	"html/template"
	"io"
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func dog(res http.ResponseWriter, req *http.Request) {
	tpl.Execute(res, nil)
}

func serve(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "."+req.URL.Path)
}

func init() {
	tpl = template.Must(template.ParseFiles("templates/dog.gohtml"))
}

var tpl *template.Template

func main() {

	// handle routes
	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dog)

	// file server
	http.HandleFunc("/assets/", serve)

	http.ListenAndServe(":8080", nil)
}

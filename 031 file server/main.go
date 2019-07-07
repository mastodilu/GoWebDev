package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request) {
	tpl.Execute(res, nil)
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

var tpl *template.Template

func main() {
	fmt.Println("Listening on port :8080")

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/index", home)

	http.ListenAndServe(":8080", nil)
}

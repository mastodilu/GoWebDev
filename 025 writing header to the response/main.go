package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Handler is my handler for incoming connections
type Handler struct{}

func (h Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("who-s-awesome", "you <3")
	tpl.ExecuteTemplate(res, "template.gohtml", nil)
}

func init() {
	tpl = template.Must(template.New("").ParseGlob("template.gohtml"))
}

var tpl *template.Template

func main() {
	fmt.Println("Server listening on port 8080")
	var Handler Handler
	http.ListenAndServe(":8080", Handler)
}

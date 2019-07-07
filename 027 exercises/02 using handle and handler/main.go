package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// Handler my http handler
type Handler struct{}

func (h Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintln(res, "<h1>HOME</h1>")
	case "/dog/":
		fmt.Fprintln(res, "<h1>DOG</h1>")
	case "/me/":
		err := req.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		name := req.FormValue("yourname")
		tpl.Execute(res, name)
	}
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

var tpl *template.Template

func main() {

	var handler Handler

	http.Handle("/", handler)

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

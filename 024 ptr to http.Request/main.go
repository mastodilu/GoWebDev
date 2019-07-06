package main

import (
	"log"
	"net/http"
	"net/url"
	"text/template"
)

// Handler is my type handler
type Handler struct{}

func (h Handler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Header        http.Header
		ContentLength int64
		M             string     // request method
		URL           *url.URL   // request url
		FormData      url.Values // contenuto del form
	}{
		Header:        request.Header,
		ContentLength: request.ContentLength,
		M:             request.Method,
		URL:           request.URL,
		FormData:      request.PostForm,
	}

	err = tpl.ExecuteTemplate(response, "index.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	tpl = template.Must(template.New("").ParseGlob("*.gohtml"))
}

var tpl *template.Template

func main() {
	var handler Handler
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}
}

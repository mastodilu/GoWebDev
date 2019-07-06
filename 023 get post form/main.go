package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type myhandler struct{}

func (m myhandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	tpl.ExecuteTemplate(responseWriter, "maintemplate.gohtml", request.Form)
}

func init() {
	tpl = template.Must(template.New("maintemplate").ParseGlob("*.gohtml"))
}

var tpl *template.Template

func main() {
	fmt.Println("Listening on port :8080")
	var myh myhandler
	http.ListenAndServe(":8080", myh)
}

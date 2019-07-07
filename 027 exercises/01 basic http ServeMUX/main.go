package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func home(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "<h1>HOME</h1>")
}

func dog(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "<h1>DOG</h1>")
}

func me(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	name := req.FormValue("yourname")
	tpl.Execute(res, name)
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

var tpl *template.Template

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

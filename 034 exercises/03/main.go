package main

import (
	"html/template"
	"net/http"
)

func serve(res http.ResponseWriter, req *http.Request) {
	path := "." + req.URL.Path
	http.ServeFile(res, req, path)
}

// func dog(res http.ResponseWriter, req *http.Request) {
// 	err := tpl.Execute(res, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

var tpl *template.Template

func main() {

	http.HandleFunc("/public/", serve)
	http.Handle("/", http.FileServer(http.Dir(".")))

	http.ListenAndServe(":8080", nil)
}

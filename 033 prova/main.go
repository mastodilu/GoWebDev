package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request) {
	tpl.Execute(res, nil)
}

func serveAsset(res http.ResponseWriter, req *http.Request) {
	assetPath := "." + req.URL.Path
	http.ServeFile(res, req, assetPath)
}

func init() {
	tpl = template.Must(template.ParseFiles("templates/tpl.gohtml"))
}

var tpl *template.Template

func main() {
	fmt.Println("Listening on port :8080")

	http.HandleFunc("/", home)
	http.HandleFunc("/assets/", serveAsset) // serves any requested asset in the folder ./assets

	http.ListenAndServe(":8080", nil)
}

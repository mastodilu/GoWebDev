package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func home(res http.ResponseWriter, req *http.Request) {
	tpl.Execute(res, nil)
}

func img(res http.ResponseWriter, req *http.Request) {
	file, err := os.Open("img.png")
	if err != nil {
		http.Error(res, "File not Found", 404)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		http.Error(res, "file not found", 404)
		return
	}

	http.ServeContent(res, req, file.Name(), stat.ModTime(), file)
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

var tpl *template.Template

func main() {
	fmt.Println("Listening on port :8080")

	http.HandleFunc("/", home)
	http.HandleFunc("/img.png", img)

	http.ListenAndServe(":8080", nil)
}

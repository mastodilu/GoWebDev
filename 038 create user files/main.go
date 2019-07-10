package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		username := r.PostFormValue("username")
		text := r.PostFormValue("text")
		path := filepath.Join("./Users/", fmt.Sprintf("%s.txt", username))
		err := ioutil.WriteFile(path, []byte(text), 0644)
		if err != nil {
			log.Println(err)
			tpl.Execute(w, nil)
			return
		}

		tpl.Execute(w, true)
		return
	}
	tpl.Execute(w, nil)
}

func init() {
	tpl = template.Must(template.ParseFiles("templates/tpl.gohtml"))
}

var tpl *template.Template

func main() {

	http.HandleFunc("/", home)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

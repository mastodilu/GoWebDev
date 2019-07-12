package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "login.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// private legge lo username e cerca il cookie con lo stesso nome.
// Se esisteva gia' allora ne prende il valore e lo incrementa di 1,
// altrimenti ne crea uno nuovo e lo inizializza a 1.
func private(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	if username == "" {
		home(w, r)
		return
	}

	loginCounter, err := r.Cookie(username)
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name:  username,
			Value: "1",
		})
	} else {
		count, err := strconv.Atoi(loginCounter.Value)
		if err != nil {
			count = 0
		}

		http.SetCookie(w, &http.Cookie{
			Name:  username,
			Value: strconv.Itoa(count + 1),
		})
	}

	err = tpl.ExecuteTemplate(w, "private.gohtml", username)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

var tpl *template.Template

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/private", private)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

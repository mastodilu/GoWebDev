package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Username string
}

func home(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Message string
		Users   []User
	}{
		Message: "Eccoti nella home",
		Users:   users,
	}
	tpl.ExecuteTemplate(w, "homepage.gohtml", data)
}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/register")

	var email, password, username string
	email = r.FormValue("email")
	password = r.FormValue("password")
	username = r.FormValue("username")
	if email == "" || password == "" || username == "" {
		log.Println("email or password are empty")
		fmt.Fprintln(w, "Error parsing your values", http.StatusNotFound)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, "Error parsing your values", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "OK")
}

func login(w http.ResponseWriter, r *http.Request) {
	var email, password string
	email = r.FormValue("email")
	password = r.FormValue("password")
	if email == "" || password == "" {
		log.Println("email or password are empty")
		fmt.Fprintln(w, "Error parsing your values", http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "ok")
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	users = append(users, User{"mastodilu"})
}

var users []User
var tpl *template.Template

func main() {

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)

	http.ListenAndServe(":8080", nil)
}

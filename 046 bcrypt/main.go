package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	db "./fakedb"

	"golang.org/x/crypto/bcrypt"
)

// UEM -
type UEM struct {
	Username string
	Email    string
	Message  string
}

func registerUser(username, email string, password []byte) error {
	hashPwd, err := bcrypt.GenerateFromPassword(password, bcrypt.MaxCost)
	if err != nil {
		return err
	}
	if err := db.Register(username, email, hashPwd); err != nil {
		return err
	}
	return nil
}

func register(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "register.gohtml", nil)
}

func save(w http.ResponseWriter, r *http.Request) {
	username := strings.TrimSpace(r.PostFormValue("username"))
	email := strings.TrimSpace(r.PostFormValue("email"))
	password := strings.TrimSpace(r.PostFormValue("password"))

	// check form data
	if username == "" || email == "" || password == "" {
		tpl.ExecuteTemplate(w, "error.gohtml", "you must enter each field")
		return
	}

	// try to register user
	if err := registerUser(username, email, []byte(password)); err != nil {
		tpl.ExecuteTemplate(w, "error.gohtml", "can't register this user")
		return
	}

	// if ok go to homepage
	http.SetCookie(w, &http.Cookie{
		Name:  "email",
		Value: email,
	})
	tpl.ExecuteTemplate(w, "okregistered.gohtml", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	cookieSession, err := r.Cookie("email")
	if err != nil {
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	email := cookieSession.Value

	user, msg, err := db.GetUsernameMessage(email)
	if err != nil {
		tpl.ExecuteTemplate(w, "error.gohtml", "non sei un utente registrato")
		return
	}

	tpl.ExecuteTemplate(w, "home.gohtml", UEM{user, email, msg})
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

var tpl *template.Template

func main() {

	http.HandleFunc("/register", register)
	http.HandleFunc("/save", save)
	http.HandleFunc("/", home)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
	fmt.Println("Hey there.")
}

// TODO cancella iscrizione utente
// TODO aggiorna messaggio utente
// TODO form 'homepage' dove punta?

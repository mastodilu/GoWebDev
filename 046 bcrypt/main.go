package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	db "./fakeDB"

	"golang.org/x/crypto/bcrypt"
)

// UEM -
type UEM struct {
	Username string
	Email    string
	Message  string
}

func registerUser(username, email string, password []byte) error {
	hashPwd, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return err
	}
	if err := db.Register(username, email, hashPwd); err != nil {
		return err
	}
	return nil
}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/register")
	tpl.ExecuteTemplate(w, "register.gohtml", nil)
}

func save(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/save")
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
	fmt.Println("here")

	// if ok go to homepage
	http.SetCookie(w, &http.Cookie{
		Name:  "email",
		Value: email,
	})
	tpl.ExecuteTemplate(w, "okregistered.gohtml", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/home")
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

func setMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/setMessage")
	if err := checkValidSession(r); err != nil {
		tpl.ExecuteTemplate(w, "error.gohtml", "non sei un utente registrato")
		return
	}

	newMessage := r.FormValue("message")
	coo, _ := r.Cookie("email")
	email := coo.Value

	if err := db.SetMessage(email, newMessage); err != nil {
		http.Error(w, "Can't set your message", http.StatusNotFound)
		return
	}

	user, msg, err := db.GetUsernameMessage(email)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusNotFound)
		return
	}
	tpl.ExecuteTemplate(w, "home.gohtml", UEM{user, email, msg})
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/deleteUser")
	if err := checkValidSession(r); err != nil {
		tpl.ExecuteTemplate(w, "error.gohtml", "non sei un utente registrato")
		return
	}

	coo, _ := r.Cookie("email")
	email := coo.Value

	err := db.DeleteUser(email)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusNotFound)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "email",
		Value:  "",
		MaxAge: -1,
	})
	tpl.ExecuteTemplate(w, "error.gohtml", "your user have been deleted from our database")
}

// checkValidSession restituisce errore se l'utente non
// ha il cookie di sessione o se non e' registrato
func checkValidSession(r *http.Request) error {
	// controlla la sessione
	coo, err := r.Cookie("email")
	if err != nil {
		return err
	}

	// controlla se e' registrato
	if !db.IsRegistered(coo.Value) {
		return fmt.Errorf("user %v is not registered", coo.Value)
	}

	return nil
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

var tpl *template.Template

func main() {

	http.HandleFunc("/register", register)
	http.HandleFunc("/save", save)
	http.HandleFunc("/", home)
	http.HandleFunc("/setMessage", setMessage)
	http.HandleFunc("/deleteUser", deleteUser)

	http.Handle("favicon.ico", http.NotFoundHandler())

	fmt.Println("listening on port 8080")
	http.ListenAndServe(":8080", nil)
	fmt.Println("Hey there.")
}

// TODO cancella iscrizione utente: serve /deleteUser

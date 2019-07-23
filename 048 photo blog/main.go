package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// User contiene alcuni dati utente
type User struct {
	Username string
	Email    string
	session  string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Message string
		Users   []User
	}{
		Message: "Eccoti nella home",
		Users:   users,
	}
	tpl.ExecuteTemplate(w, "homepage.gohtml", data)
}

func registerCheck(w http.ResponseWriter, r *http.Request) {
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

	id, err := uuid.NewV4()
	if err != nil {
		log.Println(err)
		http.Error(w, "there was an error while creating your uuid", http.StatusNotFound)
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "sessionID",
		Value: id.String(),
	})

	users = append(users, User{username, email, id.String()})

	//TODO add user to DB
	//TODO create session
	//TODO create user directory

	fmt.Fprintln(w, "OK")
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func registerPage(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "register.gohtml", nil)
}

func loginCheck(w http.ResponseWriter, r *http.Request) {
	var email, password string
	email = r.FormValue("email")
	password = r.FormValue("password")
	if email == "" || password == "" {
		log.Println("email or password are empty")
		fmt.Fprintln(w, "Error parsing your values", http.StatusNotFound)
		return
	}
	//TODO check in DB if user exists
	//TODO check valid session
	sessionID, err := uuid.NewV4()
	if err != nil {
		http.Error(w, "Your session couldn't be created", http.StatusNotFound)
		log.Println(err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "sessionID",
		Value: sessionID.String(),
	})

	users = append(users, User{"", email, sessionID.String()})
	fmt.Fprintln(w, "ok")
}

func userblog(w http.ResponseWriter, r *http.Request) {
	//TODO check valid session
	tpl.ExecuteTemplate(w, "user.gohtml", users)
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

}

var users []User
var tpl *template.Template

func main() {

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.HandleFunc("/", homePage)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/register", registerPage)

	http.HandleFunc("/user/", userblog)

	http.HandleFunc("/loginCheck", loginCheck)
	http.HandleFunc("/registerCheck", registerCheck)

	http.ListenAndServe(":8080", nil)

	//TODO implementa logout
	//TODO implementa delete account
}

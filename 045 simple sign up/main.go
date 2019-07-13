package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

func login(w http.ResponseWriter, r *http.Request) {
	sessionID, err := r.Cookie("sessionID")
	if err != nil {
		tpl.ExecuteTemplate(w, "login.gohtml", nil)
		return
	}

	if AlreadyLoggedIn(sessionID.Value) {
		http.Redirect(w, r, "/update", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.gohtml", nil)
	return
}

func getUserEmailFromRequest(r *http.Request) (UserInfo, error) {
	info := UserInfo{
		Username: r.PostFormValue("username"),
		Email:    r.PostFormValue("email"),
		Message:  "",
	}

	if info.Username == "" || info.Email == "" {
		return UserInfo{}, fmt.Errorf("username or email empty")
	}

	return info, nil
}

func home(w http.ResponseWriter, r *http.Request) {
	// no cookie: login
	// si cookie: update
	_, err := r.Cookie("sessionID")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	// TODO check user in DB
	// . not in DB: login
	// . yes in DB: update
	http.Redirect(w, r, "/update", http.StatusSeeOther)
	return
}

func save(w http.ResponseWriter, r *http.Request) {
	info, err := getUserEmailFromRequest(r)
	if err != nil {
		log.Println("/save unable to get info")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	id, err := uuid.NewV4()
	if err != nil {
		log.Println("unable to uuid.NewV4()")
		http.NotFoundHandler()
		return
	}
	err = AddUser(id.String(), info.Username, info.Email, info.Message)
	if err != nil {
		log.Println("unable to AddUser")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "sessionID",
		Value: id.String(),
	})
	info, ok := sessionDB[id.String()]
	if !ok {
		log.Println("/save key not in db")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/update", http.StatusSeeOther)
}

func update(w http.ResponseWriter, r *http.Request) {
	// check cookie
	sessionID, err := r.Cookie("sessionID")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	message := r.PostFormValue("message")
	if message == "" {
		tpl.ExecuteTemplate(w, "home.gohtml", sessionDB[sessionID.Value])
		return
	}

	// set message
	err = SetMessage(sessionID.Value, message)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	tpl.ExecuteTemplate(w, "home.gohtml", sessionDB[sessionID.Value])
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	sessionDB = make(map[string]UserInfo)
}

var tpl *template.Template

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/save", save)
	http.HandleFunc("/update", update)

	fmt.Println("listening on port :8080")
	http.ListenAndServe(":8080", nil)
}

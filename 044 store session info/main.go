package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func home(w http.ResponseWriter, r *http.Request) {
	sessionID, err := r.Cookie("SessionID")
	if err != nil {
		tpl.ExecuteTemplate(w, "login.gohtml", nil)
		return
	}

	info := struct {
		Username, Email, ID string
	}{
		sessions[sessionID.Value].Username, sessions[sessionID.Value].Email, sessionID.Value,
	}
	tpl.ExecuteTemplate(w, "login.gohtml", info)
}

func showInfo(w http.ResponseWriter, r *http.Request) {
	var id string

	// check session cookie and eventually make one
	sessionID, err := r.Cookie("SessionID")
	if err != nil {
		uid, err := uuid.NewV4()
		if err != nil {
			http.NotFoundHandler()
			log.Println(err)
			return
		}
		id = uid.String()

		http.SetCookie(w, &http.Cookie{
			Name:  "SessionID",
			Value: id,
		})
	} else {
		id = sessionID.Value
	}

	// read form
	err = r.ParseForm()
	if err != nil {
		http.NotFoundHandler()
		log.Println(err)
		return
	}

	fmt.Println("session", id)

	// save info in map
	sessions[id] = Info{
		Username: r.PostForm["username"][0],
		Email:    r.PostForm["email"][0],
	}

	info := struct {
		Username, Email, ID string
	}{
		sessions[id].Username, sessions[id].Email, id,
	}
	tpl.ExecuteTemplate(w, "show.gohtml", info)
}

func init() {
	sessions = make(map[string]Info)
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

// Info contains user specific info
type Info struct {
	Username string
	Email    string
}

var sessions map[string]Info
var tpl *template.Template

func main() {

	fmt.Println("listening on port 8080")

	http.HandleFunc("/", home)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/show", showInfo)

	http.ListenAndServe(":8080", nil)
}

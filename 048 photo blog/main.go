package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"./fakedb"
)

// UserDataImages contiene le informazioni da passare alla pagina user.gohtml
type UserDataImages struct {
	TopMessage string
	User       fakedb.User
	ImagePaths []string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Message string
		Emails  []string
	}{
		Message: "Eccoti nella home",
		Emails:  fakedb.UserList(),
	}
	tpl.ExecuteTemplate(w, "homepage.gohtml", data)
}

// registerCheck controlla il form di registrazione
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

	// add user to DB
	if err := fakedb.AddUser(email, username); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// set seesion cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "sessionID",
		Value: email,
	})

	// get current directory
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// create user assets dir
	path := filepath.Join(wd, "assets", email)
	if err := os.Mkdir(path, 0777); err != nil {
		log.Println(err)
	}

	fmt.Fprintln(w, "registered")
}

// loginPage mostra la pagina di login
func loginPage(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

// registerPage mostra la pagina di registrazione
func registerPage(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "register.gohtml", nil)
}

// loginCheck controlla il form di login
func loginCheck(w http.ResponseWriter, r *http.Request) {
	var email, password string
	email = r.FormValue("email")
	password = r.FormValue("password")
	if email == "" || password == "" {
		log.Println("email or password are empty")
		fmt.Fprintln(w, "Error parsing your values", http.StatusInternalServerError)
		return
	}
	ok := fakedb.UserExists(email)
	if !ok {
		http.Error(w, "user does not exist", http.StatusInternalServerError)
		log.Println("user does not exist")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "sessionID",
		Value: email,
	})

	fmt.Fprintln(w, "logged in")
}

func folderExists(folder string) bool {
	_, err := os.Stat(filepath.Join("users", folder))
	if err != nil {
		return !os.IsNotExist(err)
	}
	return true
}

// imageUpload controlla il form per l'inserimento di immagini nel form
func imageUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("image upload ")

	if !validSession(r) {
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	e, _ := r.Cookie("sessionID") // already checked for error
	email := e.Value

	f, h, err := r.FormFile("image")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	//TODO check file already exists
	// get current  directory
	wd, err := os.Getwd()
	if err != nil {
		log.Println("Debug 1", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	path := filepath.Join(wd, "assets", email, h.Filename)
	// create empty file
	newfile, err := os.Create(path)
	if err != nil {
		log.Println("Debug 2", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer newfile.Close()
	// write content in new empty file you just created
	if _, err = io.Copy(newfile, f); err != nil {
		log.Println("Debug 3", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := fakedb.AddImage(email, h.Filename); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// add image to user struct
	user, err := fakedb.GetUser(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	data := UserDataImages{"foto caricate!", user, user.Images}

	tpl.ExecuteTemplate(w, "user.gohtml", data)
}

// userblog mostra la pagina del blog di un utente
func userblog(w http.ResponseWriter, r *http.Request) {
	if !validSession(r) {
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}
	splitted := strings.Split(r.URL.Path, "/")
	l := len(splitted)
	if l < 2 {
		msg := fmt.Sprintf("there's a problem splitting this URL: %v", splitted)
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(msg)
		return
	}

	user, err := fakedb.GetUser(splitted[len(splitted)-1])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	data := UserDataImages{TopMessage: "un messaggio di alert a caso", User: user}
	tpl.ExecuteTemplate(w, "user.gohtml", data) // users[userIndex]
}

// checkValidSession controlla che la sessione sia attiva:
// controlla che ci sia il cookie di sessione e che sia valido
// TODO controlla la registrazione nel database
func validSession(r *http.Request) bool {
	email, err := r.Cookie("sessionID")
	if err != nil {
		return false
	}

	return fakedb.UserExists(email.Value)
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

}

var tpl *template.Template

func main() {

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/assets/", http.FileServer(http.Dir("."))) // invia il contenuto della cartella assets!!!

	http.HandleFunc("/", homePage)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/register", registerPage)

	http.HandleFunc("/user/", userblog)

	http.HandleFunc("/loginCheck", loginCheck)
	http.HandleFunc("/registerCheck", registerCheck)
	http.HandleFunc("/imageUpload", imageUpload)

	http.ListenAndServe(":8081", nil)

	//TODO implementa logout
	//TODO implementa delete account
	//TODO implementa il file server per caricare le immagini ai blog
}

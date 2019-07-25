package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	uuid "github.com/satori/go.uuid"
)

// User contiene alcuni dati utente
type User struct {
	Username string
	Email    string
	session  string
}

// Data contiene un utente e un messaggio da mostrare in cima alla pagina
type Data struct {
	TopMessage string
	User       User
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

	http.SetCookie(w, &http.Cookie{
		Name:  "email",
		Value: email,
	})

	users = append(users, User{username, email, id.String()})

	//TODO add user to DB
	//TODO create session

	//create user directory
	path := fmt.Sprintf("users/%v", email)
	if err := os.Mkdir(path, 0777); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Check the new directory", path)
	fmt.Fprintln(w, "OK")
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

func folderExists(folder string) bool {
	_, err := os.Stat(fmt.Sprintf("users/%v", folder))
	if err != nil {
		return !os.IsNotExist(err)
	}
	return true
}

// imageUpload controlla il form per l'inserimento di immagini nel form
func imageUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("image upload ")

	if err := checkValidSession(r); err != nil {
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	email, _ := r.Cookie("email") // already checked for error

	f, h, err := r.FormFile("image")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// check user folder exists
	if !folderExists(email.Value) {
		log.Printf("Folder %v does not exist\n", email.Value)
		return
	}
	//TODO store file in user dir
	var buffer []byte
	//TODO check https://golang.org/pkg/image/

	if _, err := io.ReadFull(f, buffer); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	path := fmt.Sprintf("users/%v/%v", email.Value, h.Filename)
	fmt.Println("> ", path)
	if err := ioutil.WriteFile(path, buffer, 0644); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//TODO reload user page with new content
	fmt.Printf("file %T uploaded\n", f)
	fmt.Printf("File name: %v, size: %v\n", h.Filename, h.Size)
}

// userblog mostra la pagina del blog di un utente
func userblog(w http.ResponseWriter, r *http.Request) {
	//TODO check valid session
	userIndex := 0
	data := Data{TopMessage: "un messaggio di alert a caso", User: users[userIndex]}
	fmt.Println(data)
	tpl.ExecuteTemplate(w, "user.gohtml", data) // users[userIndex]
}

// checkValidSession controlla che la sessione sia attiva:
// controlla che ci sia il cookie di sessione e che sia valido
// TODO controlla la registrazione nel database
func checkValidSession(r *http.Request) error {

	if _, err := r.Cookie("sessionID"); err != nil {
		return err
	}

	if _, err := r.Cookie("email"); err != nil {
		return err
	}
	return nil
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
	http.HandleFunc("/imageUpload", imageUpload)

	http.ListenAndServe(":8081", nil)

	//TODO implementa logout
	//TODO implementa delete account
	//TODO implementa il file server per caricare le immagini ai blog
}

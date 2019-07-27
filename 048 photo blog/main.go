package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"./fakedb"

	uuid "github.com/satori/go.uuid"
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
		Users   []fakedb.User
	}{
		Message: "Eccoti nella home",
		Users:   fakedb.UserList(),
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

	// create UUID
	id, err := uuid.NewV4()
	if err != nil {
		log.Println(err)
		http.Error(w, "there was an error while creating your uuid", http.StatusNotFound)
	}

	// add user to DB
	if err = fakedb.AddUser(email, username, id.String()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// set seesion cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "sessionID",
		Value: id.String(),
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
	user, err := fakedb.GetUser(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "sessionID",
		Value: user.SessionID,
	})

	fmt.Fprintln(w, "logged in")
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
	sessionID, _ := r.Cookie("sessionID")

	f, h, err := r.FormFile("image")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// write file
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Debug 1")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//TODO check file already exists

	path := filepath.Join(wd, "assets", email.Value, h.Filename)
	newfile, err := os.Create(path)
	if err != nil {
		fmt.Println("Debug 2")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Created", path)
	defer newfile.Close()
	if _, err = io.Copy(newfile, f); err != nil {
		fmt.Println("Debug 3")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := UserDataImages{}
	data.TopMessage = "nada di importante"
	data.User = User{"mastodilu", email.Value, sessionID.Value}
	data.ImagePaths = []string{filepath.Join("assets", email.Value, "1.jpeg")}

	tpl.ExecuteTemplate(w, "user.gohtml", data)
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

package fakedb

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"

	// used internally
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB // puntatore al database da usare per ogni query

//User contiene le informazioni degli utenti
type User struct {
	Username string
	Images   []string
}

var users = map[string]User{}

// InitDB apre la connessione al database e la testa
// se non funziona restituisce errore
func InitDB(dbURL string) error {
	var err error
	db, err = sql.Open("mysql", dbURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	return nil // :)
}

// Close termina la connessione al database
func Close() {
	db.Close()
}

// AddUser aggiunge un utente al DB
func AddUser(email, username string) error {
	query := fmt.Sprintf("INSERT INTO user (username, email) VALUES ('%s', '%s');", username, email)

	if _, err := db.Exec(query); err != nil {
		return err
	}

	return nil

	// if _, ok := users[email]; ok {
	// 	return fmt.Errorf("user already exists")
	// }

	// users[email] = User{Username: username}
	// return nil
}

// UserList restituisce l'elenco di utenti
func UserList() []string {
	query := "select email from user;"
	var users []string
	rows, err := db.Query(query)
	if err != nil {
		return []string{} // TODO return error instead
	}

	for rows.Next() {
		var s string
		if err := rows.Scan(&s); err != nil {
			return []string{}
		}
		users = append(users, s)
	}
	return users
	// var uu []string
	// for k := range users {
	// 	uu = append(uu, k)
	// }
	// return uu
}

// UserExists restituisce true se l'email esiste, altrimenti false
func UserExists(em string) bool {
	query := fmt.Sprintf("select email from user where email='%s';", em)
	row := db.QueryRow(query)
	var email string
	err := row.Scan(&email)
	if err != nil {
		log.Println(err)
		return false
	}
	fmt.Printf("%v logged in\n", email)
	return true
}

// AddImage aggiunge il path dell'immagine all'utente indicato
func AddImage(email, path string) error {
	user, ok := users[email]
	if !ok {
		return fmt.Errorf("user does not exist")
	}
	user.Images = append(user.Images, filepath.Join("assets", email, path))
	users[email] = user
	return nil
}

// RemoveUser cancella l'utente dal DB
func RemoveUser(email string) error {
	query := fmt.Sprintf("delete from user where email='%s');", email)

	if _, err := db.Exec(query); err != nil {
		return err
	}

	return nil
	//delete(users, email)
}

// GetUser restituisce l'utente corrispondente all'email, oppure errore
func GetUser(email string) (User, error) {
	user, ok := users[email]
	if !ok {
		return User{}, fmt.Errorf("user does not exist")
	}
	return user, nil
}

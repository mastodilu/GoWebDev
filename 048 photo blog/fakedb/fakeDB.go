package fakedb

import (
	"fmt"
	"path/filepath"
)

//User contiene le informazioni degli utenti
type User struct {
	Username string
	Images   []string
}

var users = map[string]User{}

// AddUser aggiunge un utente al DB
func AddUser(email, username string) error {
	//TODO implementa nel DB

	if _, ok := users[email]; ok {
		return fmt.Errorf("user already exists")
	}

	users[email] = User{Username: username}
	return nil
}

// UserList restituisce l'elenco di utenti
func UserList() []string {
	var uu []string
	for k := range users {
		uu = append(uu, k)
	}
	return uu
}

// UserExists restituisce true se l'email esiste, altrimenti false
func UserExists(em string) bool {
	_, ok := users[em]
	return ok
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
func RemoveUser(email string) {
	delete(users, email)
}

// GetUser restituisce l'utente corrispondente all'email, oppure errore
func GetUser(email string) (User, error) {
	user, ok := users[email]
	if !ok {
		return User{}, fmt.Errorf("user does not exist")
	}
	return user, nil
}

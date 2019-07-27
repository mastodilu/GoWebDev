package fakedb

import "fmt"

//User contiene le informazioni degli utenti
type User struct {
	Username, SessionID string
	Images              []string
}

var users map[string]User

// AddUser aggiunge un utente al DB
func AddUser(email, username, session string) error {
	//TODO implementa nel DB
	_, ok := users[email]
	if ok {
		return fmt.Errorf("user already exists")
	}
	users[email] = User{
		Username:  username,
		SessionID: session,
	}
	return nil
}

// AddImage aggiunge il path dell'immagine all'utente indicato
func AddImage(email, path string) error {
	user, ok := users[email]
	if !ok {
		return fmt.Errorf("user does not exist")
	}
	user.Images = append(user.Images, path)
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

// SetSession aggiorna la sessione dell'utente
func SetSession(em, se string) error {
	user, ok := users[em]
	if !ok {
		return fmt.Errorf("user does not exist")
	}
	user.SessionID = se
	users[em] = user
	return nil
}

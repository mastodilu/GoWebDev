package fakedb

import "fmt"

//User contiene le informazioni degli utenti
type User struct {
	Username, Email string
	Images          []string
}

var users map[string]User

// AddUser aggiunge un utente al DB
func AddUser(email, username, session string) error {
	//TODO implementa nel DB

	if _, ok := users[session]; ok {
		return fmt.Errorf("user already exists")
	}

	if session == "" {
		return fmt.Errorf("invalid session id")
	}

	users[session] = User{
		Username: username,
		Email:    email,
	}
	return nil
}

// UserList restituisce l'elenco di utenti
func UserList() []string {
	var uu []string
	for _, v := range users {
		uu = append(uu, v.Email)
	}
	return uu
}

// AddImage aggiunge il path dell'immagine all'utente indicato
func AddImage(session, path string) error {
	user, ok := users[session]
	if !ok {
		return fmt.Errorf("user does not exist")
	}
	user.Images = append(user.Images, path)
	users[session] = user
	return nil
}

// RemoveUser cancella l'utente dal DB
func RemoveUser(email string) {
	delete(users, email)
}

// GetUser restituisce l'utente corrispondente all'email, oppure errore
func GetUser(session string) (User, error) {
	user, ok := users[session]
	if !ok {
		return User{}, fmt.Errorf("user does not exist")
	}
	return user, nil
}

// SetSession aggiorna la sessione dell'utente
func SetSession(current, new string) error {
	user, ok := users[current]
	if !ok {
		return fmt.Errorf("user does not exist")
	}
	delete(users, current)
	users[new] = user
	return nil
}

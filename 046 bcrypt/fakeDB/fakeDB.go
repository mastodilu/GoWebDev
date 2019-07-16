package fakedb

import (
	"fmt"
)

// UserInfo contiene dei dati utente
type userInfo struct {
	Username, Message string
	hashPassword      []byte
}

// sessionDB contiene le informazioni di tutti gli utenti
var sessionDB = map[string]userInfo{}

// IsRegistered restituisce true se l'email viene trovata nel DB
func IsRegistered(email string) bool {
	_, ok := sessionDB[email]
	return ok
}

// Register aggiunge l'utente al DB.
// se è già registrato restituisce errore.
func Register(username, email string, hash []byte) error {
	if !IsRegistered(email) {
		userInfo := userInfo{
			Username:     username,
			Message:      "",
			hashPassword: hash,
		}
		sessionDB[email] = userInfo
		return nil
	}
	return fmt.Errorf("user is already registered with email %v", email)
}

// DeleteUser cancella l'utente dal DB.
// Se l'email non esiste restituisce errore
func DeleteUser(email string) error {
	if !IsRegistered(email) {
		return fmt.Errorf("user %v is not registered", email)
	}
	delete(sessionDB, email)
	return nil
}

// SetMessage aggiorna il messaggio dell'utente.
// Se l'email non esiste restituisce errore.
func SetMessage(email, msg string) error {
	if !IsRegistered(email) {
		return fmt.Errorf("user %v is not registered", email)
	}
	new, ok := sessionDB[email]
	if !ok {
		return fmt.Errorf("user %v is not registered", email)
	}
	new.Message = msg
	sessionDB[email] = new
	return nil
}

// GetUsernameMessage restituisce username e password dell'utente, oppure errore
func GetUsernameMessage(email string) (string, string, error) {
	if IsRegistered(email) {
		return sessionDB[email].Username, sessionDB[email].Message, nil
	}
	return "", "", fmt.Errorf("user is not registered")
}

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
var sessionDB map[string]userInfo

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

// GetUsernameMessage restituisce username e password dell'utente, oppure errore
func GetUsernameMessage(email string) (string, string, error) {
	if IsRegistered(email) {
		return sessionDB[email].Username, sessionDB[email].Message, nil
	}
	return "", "", fmt.Errorf("user is not registered")
}

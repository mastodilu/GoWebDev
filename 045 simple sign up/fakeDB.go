package main

import "fmt"

// UserInfo contiene dei dati utente
type UserInfo struct {
	Username, Email, Message string
}

// sessionDB contiene le informazioni di tutti gli utenti
var sessionDB map[string]UserInfo

// SetMessage setta il messaggio per per UserInfo corrispondente
func SetMessage(key, message string) error {
	info, ok := sessionDB[key]
	if !ok {
		return fmt.Errorf("session %v doesn't exist", key)
	}
	info.Message = message
	sessionDB[key] = info
	return nil
}

// AlreadyLoggedIn controlla che la sessione passata sia presente nel DB
func AlreadyLoggedIn(key string) bool {
	_, ok := sessionDB[key]
	return ok
}

// AddUser aggiunge un utente al DB
func AddUser(key, username, email, message string) error {
	info, ok := sessionDB[key]
	if !ok {
		fmt.Println("here")
		info = UserInfo{username, email, message}
		sessionDB[key] = info
		return nil
	}
	fmt.Println("there")
	return fmt.Errorf("can't add session %v because it already exists", key)
}

package main

import (
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func home(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("uuid")
	if err != nil {
		// non esiste il cookie di sessione, lo creo
		id, err := uuid.NewV4()
		if err != nil {
			http.NotFoundHandler()
			log.Fatal(err)
		}
		session := http.Cookie{
			Name:  "uuid",
			Value: id.String(),
		}
		http.SetCookie(w, &session)
		fmt.Fprintln(w, "UUID created, check your cookies")
		return
	}

	fmt.Fprintln(w, "uuid already existent, check your cookies")
}

func main() {

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", home)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

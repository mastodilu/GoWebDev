package main

import (
	"fmt"
	"log"
	"net/http"
)

// AkunaMatata tipo del tutto casuale che implementa un http handler da passare a ListenAndServe
type AkunaMatata bool

func (a AkunaMatata) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ServeHTTP method")
	fmt.Fprintf(w, "%s", "B00BS  ")
}

func main() {
	var akmt AkunaMatata
	fmt.Println("Listening on port :8080")
	err := http.ListenAndServe(":8080", akmt)
	if err != nil {
		log.Fatal(err)
	}
}

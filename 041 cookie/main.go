package main

import (
	"fmt"
	"net/http"
)

func setMyCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "supah-secret",
		Value: "Culo chi legge",
	})
	fmt.Fprintln(w, "Controlla i cookie.")
}

func main() {
	fmt.Println("Listening on port :8080")

	http.HandleFunc("/", setMyCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

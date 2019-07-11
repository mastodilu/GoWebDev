package main

import (
	"fmt"
	"net/http"
)

func setMyCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "myCookie",
		Value: "Ciaone",
	})
	fmt.Fprintln(w, "Controlla i cookie.")
}

func readCookie(w http.ResponseWriter, r *http.Request) {
	for i, c := range r.Cookies() {
		fmt.Fprintln(w, i, *c)
	}

	c, err := r.Cookie("myCookie")
	if err != nil {
		fmt.Println("No cookie named 'myCookie'")
	}
	fmt.Fprintln(w, "Found 'myCookie': ", *c)
}

func main() {
	fmt.Println("Listening on port :8080")

	http.HandleFunc("/", setMyCookie)
	http.HandleFunc("/read", readCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

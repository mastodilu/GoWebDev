package main

import (
	"fmt"
	"net/http"
)

// setMyCookie scrive un cookie chiamato 'myCookie'
func setMyCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "myCookie",
		Value: "Ciaone",
	})
	fmt.Fprintln(w, "Controlla i cookie.")
}

// readCookie legge tutti i cookie
func readCookie(w http.ResponseWriter, r *http.Request) {
	for i, c := range r.Cookies() {
		fmt.Fprintln(w, i, *c)
	}

	fmt.Fprintln(w, "######### ######### #########")

	c, err := r.Cookie("myCookie")
	if err != nil {
		fmt.Println("No cookie named 'myCookie'")
	}
	fmt.Fprintln(w, "Found 'myCookie': ", *c)
}

// writeAll scrive un botto di cookie con un ciclo
func writeAll(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 5; i++ {
		http.SetCookie(w, &http.Cookie{
			Name:  fmt.Sprintf("Cookie-%v", i+1),
			Value: fmt.Sprintf("Questo cookie vale %v", i+1),
		})
	}
}

// deleteCookie cancella il cookie di nome 'myCookie'
func deleteCookie(w http.ResponseWriter, r *http.Request) {
	ck, _ := r.Cookie("myCookie")
	if ck != nil {
		ck.MaxAge = -1 // deletes the cookie!
		http.SetCookie(w, ck)
		fmt.Fprintln(w, "'myCookie' deleted")
	}
}

func main() {
	fmt.Println("Saluti da Linux Mint")
	fmt.Println("Listening on port :8080")

	http.HandleFunc("/", setMyCookie)
	http.HandleFunc("/read", readCookie)
	http.HandleFunc("/write-all", writeAll)
	http.HandleFunc("/delete", deleteCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

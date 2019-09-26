package main

import (
	"context"
	"fmt"
	"net/http"
)

// User -
type User struct {
	Username, Email string
}

// Key -
type Key string // 💥 serve a fare da wrapper al tipo string che non viene accettato da context

func register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// var usernameKey Key = "username" // non usare tipi base come chiavi, ma altro, tipo 'type Key string'
	// var emailKey Key = "email"       // si evitano conflitti che boh non ho capito ma va beh funziona lol ok ciao :)
	ctx = context.WithValue(ctx, Key("username"), "mastodilu")   // 💥
	ctx = context.WithValue(ctx, Key("email"), "text@email.com") // 💥
	if addUserToDB(ctx) {
		fmt.Fprintf(w, "%v", "Ok")
	} else {
		fmt.Fprintf(w, "%v", "Not ok")
	}
}

func printContext(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println(ctx)
	fmt.Fprintf(w, "%v%v", "<h1>Context</h1>", ctx)
}

func addUserToDB(ctx context.Context) bool {
	username := ctx.Value(Key("username")).(string)
	email := ctx.Value(Key("email")).(string)
	if username != "" && email != "" {
		fmt.Println("Read from context:", username, email)
		return true
	} else {
		fmt.Println("Found nothing in context")
		return false
	}
}

var db = make(map[string]User)

func main() {
	http.HandleFunc("/register", register)
	http.HandleFunc("/context", printContext)
	http.ListenAndServe(":8080", nil)
}

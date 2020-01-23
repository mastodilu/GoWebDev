package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mastodilu/GoWebDev/053_mongoDB/02_JSON/models"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	home := `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<meta http-equiv="X-UA-Compatible" content="ie=edge">
	<title>Home</title>
</head>
<body>
	<h1>Click the link to get a user</h1>
	<a href="/user/1234566789">see user 1234566789</a>
</body>
</html>`
	fmt.Fprintln(w, home)
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "matteo",
		Gender: "male",
		Age:    25,
		ID:     "0",
	}
	j, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s\n", j)
}

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	fmt.Println("server started on port 8080")
	http.ListenAndServe("localhost:8080", r)
}

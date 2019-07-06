package main

import (
	"fmt"
	"net/http"
)

func cat(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<h1>Handles %v</h1>\n%v\n", "/catOnly", req.URL.Path)
}

func dog(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<h1>Handles %v</h1>\n%v\n", "/dog/", req.URL.Path)
}

func home(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<h1>Handles %v</h1>\n%v\n", "/", req.URL.Path)
}

func main() {
	fmt.Println("server listening on port :8080")

	http.HandleFunc("/", home)
	http.HandleFunc("/catOnly", cat)
	http.HandleFunc("/dog/", dog)

	http.ListenAndServe(":8080", nil)
}

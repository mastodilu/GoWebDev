package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello golang from docker")
}

func main() {
	fmt.Println("Listening on port :8081")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8081", nil)
}

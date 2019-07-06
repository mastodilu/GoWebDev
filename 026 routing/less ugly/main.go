package main

import (
	"fmt"
	"net/http"
)

// CatHandler is my specific handler type
type CatHandler struct{}

// DogHandler is my specific handler type
type DogHandler struct{}

// Handler is my specific handler type
type Handler struct{}

func (h CatHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<h1>Handles %v</h1>\n%v\n", "/catOnly", req.URL.Path)
}

func (h DogHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<h1>Handles %v</h1>\n%v\n", "/dog/", req.URL.Path)
}

func (h Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<h1>Handles %v</h1>\n%v\n", "/", req.URL.Path)
}

func main() {
	fmt.Println("server listening on port :8080")

	// handlers
	cat := CatHandler{}
	dog := DogHandler{}
	boringHandler := Handler{}

	mux := http.NewServeMux()
	mux.Handle("/", boringHandler)
	mux.Handle("/catOnly", cat) // only route /cat
	mux.Handle("/dog/", dog)    // route /dog and /dog/every/other

	http.ListenAndServe(":8080", mux)
}

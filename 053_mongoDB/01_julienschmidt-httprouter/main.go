package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"../02 JSON/models"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hi\n")
}

func main() {
	r := httprouter.New()
	r.GET("/", index)
	fmt.Println("server started on port 80800")
	http.ListenAndServe("localhost:8080", r)
}

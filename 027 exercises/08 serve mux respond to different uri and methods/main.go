package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"text/template"
)

type Data struct {
	Username string
	URL      *url.URL
	Method   string
}

func indexGET(res http.ResponseWriter, req *http.Request) {
	data := Data{
		Username: "",
		URL:      req.URL,
		Method:   req.Method,
	}
	tpl.Execute(res, data)
}

func indexPOST(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	data := Data{
		Username: req.FormValue("username"),
		URL:      req.URL,
		Method:   req.Method,
	}
	fmt.Println(data)
	tpl.Execute(res, data)
}

func index(res http.ResponseWriter, req *http.Request) {
	switch strings.ToLower(req.Method) {
	case "get":
		fmt.Println("index get")
		indexGET(res, req)
	case "post":
		fmt.Println("index post")
		indexPOST(res, req)
	default:
		fmt.Println("I don't know what to do")
	}
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

var tpl *template.Template

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/index/", index)

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"os"
	"text/template"
)

func init() {
	tpl = template.Must(template.New("maintemplate.gohtml").ParseFiles("templates/maintemplate.gohtml"))
}

var tpl *template.Template

func main() {
	xss := `<script>alert("wela")</script>`
	err := tpl.Execute(os.Stdout, xss)
	if err != nil {
		panic(err)
	}
}

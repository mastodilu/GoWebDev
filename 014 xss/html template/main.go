package main

import (
	"html/template"
	"os"
)

func init() {
	tpl = template.Must(template.New("maintemplate.gohtml").ParseFiles("templates/maintemplate.gohtml"))
}

var tpl *template.Template

func main() {
	data := `<script>alert("wela")</script>`
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

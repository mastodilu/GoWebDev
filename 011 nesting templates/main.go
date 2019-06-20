package main

import (
	"mylib"
	"os"
	"text/template"
)

func init() {
	tpl = template.Must(template.New("").ParseGlob("templates/*.gohtml"))
}

var tpl *template.Template

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "mainTemplate.gohtml", nil)
	mylib.IfErrThenLogFatal(err, "can't execute mainTemplate.gohtml")
}

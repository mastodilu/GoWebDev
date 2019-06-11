package main

import (
	"html/template"
	"mylib"
	"os"
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

var tpl *template.Template

func main() {
	myData := "wela"
	err := tpl.ExecuteTemplate(os.Stdout, "one.gohtml", myData)
	mylib.IfErrThenLogFatal(err, "Can't execute template")
}

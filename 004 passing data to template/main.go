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

	err := tpl.ExecuteTemplate(os.Stdout, "one.gohtml", "wela")
	mylib.IfErrThenLogFatal(err, "Can't execute template")
}

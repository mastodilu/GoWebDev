package main

import (
	"mylib"
	"os"
	"text/template"
)

func subt1(n int) int {
	return n - 1
}

func add1(n int) int {
	return n + 1
}

func init() {
	funcmap = template.FuncMap{
		"add1":  add1,
		"subt1": subt1,
	}
	tpl = template.Must(template.New("").Funcs(funcmap).ParseGlob("templates/*.gohtml"))
}

var tpl *template.Template
var funcmap template.FuncMap

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "pipeline.gohtml", 0)
	mylib.IfErrThenLogFatal(err, "can't execute one.gohtml template")
}

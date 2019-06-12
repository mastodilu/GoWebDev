package main

import (
	"html/template"
	"mylib"
	"os"
)

func init() {
	tpl = template.Must(template.New("").ParseGlob("templates/*.gohtml"))
}

// var funcmap template.FuncMap
var tpl *template.Template

func main() {
	myData := struct {
		Name string
		Data []string
	}{
		Name: "matteo",
		Data: []string{"zero", "one", "two", "three"},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "globfunc.gohtml", myData)
	mylib.IfErrThenLogFatal(err, "can't execute globafunc.gohtml template")
}

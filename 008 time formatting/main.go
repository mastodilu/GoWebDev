package main

import (
	"mylib"
	"os"
	"text/template"
	"time"
)

func formatTime(t time.Time) string {
	myFormat := "06/01/02 03:04 PM MST"
	return t.Format(myFormat)
}

func init() {
	funcmap = template.FuncMap{
		"formatTime": formatTime,
	}
	tpl = template.Must(template.New("").Funcs(funcmap).ParseGlob("templates/*.gohtml"))
}

var tpl *template.Template
var funcmap template.FuncMap

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "one.gohtml", time.Now())
	mylib.IfErrThenLogFatal(err, "can't execute one.gohtml template")
}

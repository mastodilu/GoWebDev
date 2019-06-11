package main

import (
	"fmt"
	"html/template"
	"mylib"
	"os"
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

var tpl *template.Template

func main() {
	mySlice := []string{"one", "two", "three"}
	err := tpl.ExecuteTemplate(os.Stdout, "slice.gohtml", mySlice)
	mylib.IfErrThenLogFatal(err, "Can't execute template")

	fmt.Println("\n------------")

	myMap := map[string]string{
		"uno": "one",
		"due": "two",
		"tre": "three",
	}
	err = tpl.ExecuteTemplate(os.Stdout, "map.gohtml", myMap)
	mylib.IfErrThenLogFatal(err, "Can't execute template")
}

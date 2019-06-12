package main

import (
	"fmt"
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
		Name      string
		Data      []string
		ShortData []int
	}{
		Name:      "matteo",
		Data:      []string{"zero", "one", "two", "three"},
		ShortData: []int{3, 6},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "globfunc.gohtml", myData)
	mylib.IfErrThenLogFatal(err, "can't execute globafunc.gohtml template")

	fmt.Println("------------------------")

	err = tpl.ExecuteTemplate(os.Stdout, "if.gohtml", myData)
	mylib.IfErrThenLogFatal(err, "can't execute if.gohtml template")
}

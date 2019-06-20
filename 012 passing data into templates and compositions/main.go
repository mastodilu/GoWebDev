package main

import (
	"mylib"
	"os"
	"text/template"
)

func init() {
	tpl = template.Must(template.New("").ParseGlob("templates/*.gohtml"))
	out, _ = os.Create("out.html")
}

type People struct {
	Name    string
	Surname string
	Classes []ClassScore
}

type ClassScore struct {
	Class string
	Score uint
}

var tpl *template.Template
var out *os.File

func main() {
	people := []People{
		{
			Name: "matteo",
			Surname: "dilu",
			Classes: []ClassScore{
				{"Maths", 9},
				{"Science", 8},
				{"English", 7},
			},
		},
		{
			Name: "giovanno",
			Surname: "pagliaccio",
			Classes: []ClassScore{
				{"Maths", 8},
				{"Science", 10},
				{"English", 9},
			},
		},
		{
			Name: "bimbum",
			Surname: "bam",
			Classes:[]ClassScore{
				{"Maths", 6},
				{"Science", 9},
				{"English", 9},
			},
		},
	}
	
	err := tpl.ExecuteTemplate(out, "main.gohtml", people)
	mylib.IfErrThenPanic(err, "can't execute main.gohtml template")
}

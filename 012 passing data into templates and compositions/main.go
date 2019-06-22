package main

import (
	"os"
	"text/template"
)

func init() {
	tpl = template.Must(template.New("").ParseGlob("templates/*.gohtml"))
	out, _ = os.Create("out.html")
}

// People è un tipo che contiene nome, cognome e slice di ClassScore
type People struct {
	Name    string
	Surname string
	Classes []ClassScore
}

// ClassScore è una struct che contiene una coppia classe e voto
type ClassScore struct {
	Class string
	Score uint
}

// SayHi stampa un messaggio per il tipo People
func (p People) SayHi() string {
	return "hi people!"
}

// SayHi stampa un messaggio per il tipo ClassScore
func (c *ClassScore) SayHi() string {
	return "hi class!"
}

var tpl *template.Template
var out *os.File

func main() {
	people := []People{
		{
			Name:    "matteo",
			Surname: "dilu",
			Classes: []ClassScore{
				{"Maths", 9},
				{"Science", 8},
				{"English", 7},
			},
		},
		{
			Name:    "giovanno",
			Surname: "pagliaccio",
			Classes: []ClassScore{
				{"Maths", 8},
				{"Science", 10},
				{"English", 9},
			},
		},
		{
			Name:    "bimbum",
			Surname: "bam",
			Classes: []ClassScore{
				{"Maths", 6},
				{"Science", 9},
				{"English", 9},
			},
		},
	}

	err := tpl.ExecuteTemplate(out, "main.gohtml", people)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"os"
	"text/template"
)

type Hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

type CaliforniaHotels []Hotel

func init() {
	tpl = template.Must(template.ParseFiles("template/mainTemplate.gohtml"))
}

var tpl *template.Template

func main() {
	ch := CaliforniaHotels{
		Hotel{
			Name:    "HotelCaliforniaaah",
			Address: "123 l'hotel è di lè",
			City:    "Città di california",
			Zip:     "zap zap",
			Region:  "South",
		},
		Hotel{
			Name:    "noname",
			Address: "noaddress",
			City:    "cityh",
			Zip:     "zum zum ZUM",
			Region:  "si",
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "mainTemplate.gohtml", ch)
	if err != nil {
		panic(err)
	}
}

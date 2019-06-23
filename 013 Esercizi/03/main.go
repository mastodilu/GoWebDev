package main

import (
	"os"
	"text/template"
)

// Restaurant è una struct che rappresenta il menu di un ristorante
type Restaurant struct {
	Breakfasts []Meal
	Lunchs     []Meal
	Dinners    []Meal
}

// Meal rappresenta un pasto
type Meal struct {
	Name, Description string
	Price             float32
}

// add1 incrementa di uno il valore passato
func add1(n int) int {
	return n + 1
}

func init() {
	funcmap := template.FuncMap{
		"add1": add1,
	}
	tpl = template.Must(template.New("maintemplate.gohtml").Funcs(funcmap).ParseFiles("templates/maintemplate.gohtml"))
}

var tpl *template.Template

func main() {
	restaurants := []Restaurant{
		Restaurant{
			Breakfasts: []Meal{
				{
					Name:        "Biscotti",
					Description: "Biscotti",
					Price:       4.50,
				}, {
					Name:        "Fette biscottate",
					Description: "Delle brutte fette biscottate",
					Price:       1.20,
				},
				{
					Name:        "Frutta",
					Description: "Frutta",
					Price:       3.50,
				},
			},
			Lunchs: []Meal{
				{
					Name:        "Spaghetti",
					Description: "Spaghetti freschi",
					Price:       12.50,
				}, {
					Name:        "Pizza",
					Description: "Pizza fresca",
					Price:       10.00,
				},
				{
					Name:        "Bistecca",
					Description: "Fiorentina",
					Price:       25.50,
				},
			},
			Dinners: []Meal{
				{
					Name:        "Minestrone",
					Description: "Minestrone di verdure",
					Price:       8.50,
				}, {
					Name:        "Hotdog",
					Description: "Hotdog",
					Price:       7.20,
				},
				{
					Name:        "Pesce",
					Description: "Fritto misto",
					Price:       30.50,
				},
			},
		},
		Restaurant{
			Breakfasts: []Meal{
				{
					Name:        "Uova",
					Description: "Uova",
					Price:       4.50,
				}, {
					Name:        "Fette biscottate",
					Description: "Delle brutte fette biscottate",
					Price:       1.20,
				},
				{
					Name:        "Wurstel",
					Description: "Wurstel",
					Price:       3.50,
				},
			},
			Lunchs: []Meal{
				{
					Name:        "Spaghetti and meatballs",
					Description: "Spaghetti and meatballs",
					Price:       12.50,
				}, {
					Name:        "Pizza",
					Description: "Pizza fresca",
					Price:       10.00,
				},
				{
					Name:        "Bistecca",
					Description: "Fiorentina",
					Price:       25.50,
				},
			},
			Dinners: []Meal{
				{
					Name:        "Minestrone",
					Description: "Minestrone di verdure",
					Price:       8.50,
				}, {
					Name:        "Hamburger",
					Description: "hamburger",
					Price:       7.20,
				},
				{
					Name:        "Pesce",
					Description: "Orata",
					Price:       30.50,
				},
			},
		},
		Restaurant{
			Breakfasts: []Meal{
				{
					Name:        "Biscotti",
					Description: "Biscotti",
					Price:       4.50,
				}, {
					Name:        "Fette biscottate",
					Description: "Delle brutte fette biscottate",
					Price:       1.20,
				},
				{
					Name:        "Frutta",
					Description: "Frutta",
					Price:       3.50,
				},
			},
			Lunchs: []Meal{
				{
					Name:        "Spaghetti",
					Description: "Spaghetti freschi",
					Price:       12.50,
				}, {
					Name:        "Pizza",
					Description: "Pizza fresca",
					Price:       10.00,
				},
				{
					Name:        "Bistecca",
					Description: "Fiorentina",
					Price:       25.50,
				},
			},
			Dinners: []Meal{
				{
					Name:        "Minestrone",
					Description: "Minestrone di verdure",
					Price:       8.50,
				}, {
					Name:        "Hotdog",
					Description: "Hotdog",
					Price:       7.20,
				},
				{
					Name:        "Pesce",
					Description: "Fritto misto",
					Price:       30.50,
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, restaurants)
	if err != nil {
		panic(err)
	}
}

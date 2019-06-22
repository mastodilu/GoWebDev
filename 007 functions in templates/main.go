package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

var funcmap = template.FuncMap{
	"sayhi":      sayHi,
	"sum":        sum,
	"toupper":    strings.ToUpper,
	"firstthree": firstThree,
}

func init() {
	// tpl := template.New("functions.gohtml")
	// tpl = tpm.Funcs(funcmap)
	// tpl, err := template.ParseFiles(".templates/functions.gohtml")
	// tpl = template.Must(tpl, err)
	tpl = template.Must(template.New("").Funcs(funcmap).ParseFiles("templates/functions.gohtml"))
}

func sum(n1, n2 int) int {
	return n1 + n2
}

func sayHi() string {
	return "hi from sayHi() function"
}

func firstThree(s string) (string, error) {
	r := []rune(strings.TrimSpace(s))
	if len(r) < 3 {
		return "", fmt.Errorf("string '%s' is too short", s)
	}
	return string(r[:3]), nil
}

// funzioni da passare al template

var tpl *template.Template

func main() {
	fileout, err := os.Create("fileout.html")
	if err != nil {
		panic("can't create fileout.html")
	}
	defer fileout.Close()

	myData := []string{"ciao asd asd", "miao asd asd", "bau asd asd"}

	err = tpl.ExecuteTemplate(os.Stdout, "functions.gohtml", myData)
	if err != nil {
		panic("can't execute functions.html template")
	}
}

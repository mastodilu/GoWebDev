package main

import (
	"fmt"
	"mylib"
	"os"
	"text/template"
)

// stampa i nomi dei template parsati e salvati
func printAllTemplates(t *template.Template) {
	for _, tpl := range t.Templates() {
		fmt.Println(tpl.Name())
	}
}

func main() {

	const parseWholeFolder = "templates/*"
	const parseOnlyHTML = "templates/*.gohtml"
	tpl, err := template.ParseGlob(parseWholeFolder)
	mylib.IfErrThenPanic(err, "err in ParseGlob")
	fmt.Printf("%T, %v\n", tpl, tpl)

	// stampa l'elenco dei template
	printAllTemplates(tpl)

	tpl.ExecuteTemplate(os.Stdout, "sample.gnegnegne", nil)

}

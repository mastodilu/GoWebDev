package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

// init è una funzione "reservata" che viene sempre chiamata implicitamente avviando un progetto
func init() {
	const parseWholeFolder = "templates/*"
	const parseOnlyHTML = "templates/*.gohtml"

	//controlla che il template parsato non sia null
	tpl = template.Must(template.ParseGlob(parseWholeFolder))
}

func main() {

	fmt.Printf("%T, %v\n", tpl, tpl)

	tpl.ExecuteTemplate(os.Stdout, "sample.gnegnegne", nil)

}

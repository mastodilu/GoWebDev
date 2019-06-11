package main

import (
	"fmt"
	"mylib"
	"text/template"
)

func main() {

	const parseWholeFolder = "templates/*"
	const parseOnlyHTML = "templates/*.gohtml"
	tpl, err := template.ParseGlob(parseWholeFolder)
	mylib.IfErrThenPanic(err, "err in ParseGlob")
	fmt.Printf("%T, %v\n", tpl, tpl)

}

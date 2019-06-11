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
	//slice
	mySlice := []string{"one", "two", "three"}
	err := tpl.ExecuteTemplate(os.Stdout, "slice.gohtml", mySlice)
	mylib.IfErrThenLogFatal(err, "can't execute slice template")

	fmt.Println("\n------------")

	//map
	myMap := map[string]string{
		"uno": "one",
		"due": "two",
		"tre": "three",
	}
	err = tpl.ExecuteTemplate(os.Stdout, "map.gohtml", myMap)
	mylib.IfErrThenLogFatal(err, "can't execute map template")

	fmt.Println("\n------------")

	//struct (struct anonima)
	mystruct := struct {
		Name string
		Age  int
	}{
		Name: "matteo",
		Age:  25,
	}
	err = tpl.ExecuteTemplate(os.Stdout, "struct.gohtml", mystruct)
	mylib.IfErrThenLogFatal(err, "can't execute struct template")

	fmt.Println("\n------------")

	//slice di struct (anonime)
	slicestruct := []struct {
		Name string
		Age  int
	}{
		{"matteo", 25},
		{"alberto", 28},
		{"pollo", 12},
	}
	tpl.ExecuteTemplate(os.Stdout, "sliceStruct.gohtml", slicestruct)
	mylib.IfErrThenLogFatal(err, "can't execute sliceStruct template")
}

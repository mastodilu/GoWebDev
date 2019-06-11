package main

import (
	"mylib"
	"os"
	"text/template"
)

/*
	- i template sono degli scheletri che rappresentano del contenuto html

	- i file htlm usati da Go possono avere il formato che si vuole, ma è buona prassi usare '.gohtml'

	- per fare il parsing si usa il package 'template.ParseFiles' che restituisce un 'pointer a template'
		un pointer a template è come un container in cui i file sono parsificati e tenuti
*/

func main() {
	templatesFolder := "./pages/templates"

	//mytemplate è un puntatore a template, un contenitore di tutti i template che gli ho fatto parsificare
	mytemplate, err := template.ParseFiles(templatesFolder + "/htmlpage.gohtml")
	mylib.IfErrThenLogFatal(err, "can't parse files")

	//crea il file homepage.html in cui scrivere il template
	homepage, err := os.Create("./pages/out/homepage.html")
	mylib.IfErrThenLogFatal(err, "can't create homepage.html")
	defer homepage.Close()

	// scrive il template nel file homepage.html
	err = mytemplate.Execute(homepage, nil)
	mylib.IfErrThenLogFatal(err, "can't write to output file")

	//parsing di molti template
	mytemplate, err = template.ParseFiles(templatesFolder+"/one.gohtml", templatesFolder+"/two.gohtml")
	mylib.IfErrThenLogFatal(err, "can't parse all those ugly templates")

	//per scrivere su io.Reader uno specifico template quando ce ne sono tanti si usa ExecuteTemplate
	//	attento al nome specificato perchè non corrisponde al path, ma al base name
	err = mytemplate.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	mylib.IfErrThenLogFatal(err, "can't write one.txt to os.Stdout")

}

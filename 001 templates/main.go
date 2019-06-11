package main

import (
	"log"
	"os"
	"text/template"
)

/*
	- i template sono degli scheletri che rappresentano del contenuto html

	- i file htlm usati da Go possono avere il formato che si vuole, ma è buona prassi usare '.gohtml'

	- per fare il parsing si usa il package 'template.ParseFiles' che restituisce un 'pointer a template'
		un pointer a template è come un container in cui i file sono parsificati e tenuti
*/

func ifErrThenLogFatal(err error, msg ...string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func main() {
	templatesFolder := "./pages/templates"

	//mytemplate è un puntatore a template, un contenitore di tutti i template che gli ho fatto parsificare
	mytemplate, err := template.ParseFiles(templatesFolder + "/htmlpage.gohtml")
	ifErrThenLogFatal(err, "can't parse files")

	//crea il file homepage.html in cui scrivere il template
	homepage, err := os.Create("./pages/out/homepage.html")
	ifErrThenLogFatal(err, "can't create homepage.html")
	defer homepage.Close()

	// execute prende un io.Writer e dei dati (o nil)
	// questo stampa nel terminale il template parsificato
	//err = mytemplate.Execute(os.Stdout, nil)
	//ifErrThenLogFatal(err, "can't execute templates")

	// scrive il template nel file homepage.html
	err = mytemplate.Execute(homepage, nil)
	ifErrThenLogFatal(err, "can't write to output file")

	//parsing di molti template
	mytemplate, err = template.ParseFiles(templatesFolder+"/one.gohtml", templatesFolder+"/two.gohtml")
	ifErrThenLogFatal(err, "can't parse all those ugly templates")

	// apro i file in cui scrivere i template
	// one, err := os.Open("./pages/one.txt")
	//ifErrThenLogFatal(err, "can't create one.txt")
	//two, err := os.Open("./pages/two.py")
	//ifErrThenLogFatal(err, "can't create two.html")
	//three, err := os.Open("./three.shish")
	//ifErrThenLogFatal(err, "can't create three.yeah")

	//per scrivere su io.Reader uno specifico template quando ce ne sono tanti si usa ExecuteTemplate
	err = mytemplate.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	ifErrThenLogFatal(err, "can't write one.txt to os.Stdout")

}

package main

import (
	"text/template"
	"log"
	"os"
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
	filename := "./pages/htmlpage.gohtml"

	//mytemplate è un puntatore a template, un contenitore di tutti i template che gli ho fatto parsificare
	mytemplate, err := template.ParseFiles(filename)
	ifErrThenLogFatal(err, "can't parse files")

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
}

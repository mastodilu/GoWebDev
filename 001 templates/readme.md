# Readme

## Templates

[golang.org/pkg/text/template/](https://golang.org/pkg/text/template/)

Template: in pratica rappresenta lo scheletro di un sito web

## Scrivi un template in un io.Reader (file o os.Stdout)

## 1. Template singolo

```go
// Restituisce un puntatore a template che contiene il template parsificato
mytemplate  , err := template.ParseFiles(templateName)
ifErrThenLogFatal(err, "can't parse files")

//crea il file homepage.html in cui scrivere il template
homepage, err := os.Create("./pages/out/homepage.html")
ifErrThenLogFatal(err, "can't create homepage.html")
defer homepage.Close()

// scrive il template nel file homepage.html
err = mytemplate.Execute(homepage, nil)
ifErrThenLogFatal(err, "can't write to output file")
```

E' possibile stampare anche nel terminale usando os.Stdout invece di un file.

```go
// execute prende un io.Writer e dei dati (o nil)
// questo stampa nel terminale il template parsificato
err = mytemplate.Execute(os.Stdout, nil)
ifErrThenLogFatal(err, "can't execute templates")
```

## 2. Template multipli

```go
//parsing di molti template
mytemplate, err = template.ParseFiles("one.txt", "two.html", "three.yeah")
ifErrThenLogFatal(err, "can't parse all those ugly templates")

//ExecuteTemplate permette di selezionare quale template scrivere quando ce ne sono tanti
err = mytemplate.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
ifErrThenLogFatal(err, "can't write one.txt to os.Stdout")
```

Il nome del template specificato è il base name del file parsato e non path/../basename

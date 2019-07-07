# readme

## Serve files with `http.FileServer`

FileServer permette di inviare file e intere directory.

La funzione `http.FileServer` ha questa firma:

```Go
func FileServer(root FileSystem) Handler
```

e restituisce un `http.Handler`.

Con il codice

```Go
http.Handle("/", http.FileServer(http.Dir(".")))
```

stiamo dicendo di usare FileServer come handler per il percorso "/".

Questo comando ha un effetto, invia tutti i file contenuti nella cartella `http.Dir(".")` (la cartella corrente) alla pagina `localhost:8080`:

```html
<pre>
<a href="README.md">README.md</a>
<a href="img.png">img.png</a>
<a href="main.go">main.go</a>
<a href="tpl.gohtml">tpl.gohtml</a>
</pre>
```

Cliccando sul [tpl.gothml](tpl.gothml) si ottiene il caricamento della pagina che contiene l'immagine che volevamo mostrare.

### Script

```Go
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request) {
	tpl.Execute(res, nil)
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

var tpl *template.Template

func main() {
	fmt.Println("Listening on port :8080")

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/index", home)

	http.ListenAndServe(":8080", nil)
}
```

### Template

```Gohtml
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=>, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    <h1>BatGopher</h1>
    <img src="img.png" alt="BatGopher">
</body>
</html>
```

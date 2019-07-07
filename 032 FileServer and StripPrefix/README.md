# readme

Per inviare il contenuto di un'intera cartella in modo sicuro si usa `http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))`.

Questo codice fa queste cose:

1. `http.Handle` prende come parametri un `path` e un `http.Handler`
2. `http.StripPrefix(..)` prende come parametri un `path` e un `http.Handler`
3. `http.FileServer(..)` invia il contenuto della cartella specificata in modo sicuro e anti hacker

il codice in pratica gestisce le richieste che cominciano con `/resources/`, toglie  `/resources` e invia il contenuto della cartella `/assets`.

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
	tpl = template.Must(template.ParseFiles("templates/tpl.gohtml"))
}

var tpl *template.Template

func main() {
	fmt.Println("Listening on port :8080")

	http.HandleFunc("/", home)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))

	http.ListenAndServe(":8080", nil)
}
```
# Readme

## julienschmidt httprouter

[julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)

```Go
package main

import (
    "fmt"
    "net/http"

    "github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Hi\n")
}

func main() {
    r := httprouter.New()
    r.GET("/", index)
    fmt.Println("server started on port 80800")
    http.ListenAndServe("localhost:8080", r)
}
```

Si crea un nuovo httprouter

```Go
r := httprouter.New()
```

dato come parametro al metodo `http.ListenAndServe()`.

Per ascoltare una richiesta `GET` si usa `r.GET("/", index)`, dove index è una funzione che vuole come parametri:

1. risposta
2. richiesta
3. parametri dell'`httprouter` (nell'esempio il parametro non ci serve quindi lo nominiamo `_`)

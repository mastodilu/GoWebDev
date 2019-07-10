# http.NotFoundHandler()

La funzione `http.NotFoundHandler()` è così implementata:

```Go
// NotFoundHandler returns a simple request handler
// that replies to each request with a ``404 page not found'' reply.
func NotFoundHandler() Handler { return HandlerFunc(NotFound) }
```

e permette di rispondere `404 page not found`.

Ad esempio:

```Go
http.Handle("/random/url", http.NotFoundHandler())
```

---

# favicon.ico

E' un'icona standard richiesta dai browser in modo automatico come icona nella tab della pagina web visualizzata.

Per inviarla è possibile fare in questo modo:

```Go
func favicon(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "assets/favicon.ico")
}

http.HandleFunc("/favicon.ico", favicon)

// oppure così se non è disponibile
http.Handle("/favicon.ico", http.NotFoundHandler())
```

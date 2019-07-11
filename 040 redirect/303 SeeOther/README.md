# 301 Moved Permanently

La redirect **301** mantiene lo stesse metodo della request originale.

Si effettua in questi due modi alternativi:

```Go
func redirect1(w http.ResponseWriter, r *http.Request) {
    http.Re direct(w, r, "/", http.StatusSeeOther)
}

func redirect2(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Location", "/")
    w.WriteHeader(http.StatusSeeOther)
}
```

```Go
package main

import (
    "fmt"
    "net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprintf(w, "<h1>%v</h1>\n", "Home")
}

func redirect1(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/", http.StatusSeeOther) // 💥
}

func redirect2(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Location", "/")    // 💥
    w.WriteHeader(http.StatusSeeOther) // 💥
}

func main() {
    fmt.Println("Listening on port :8080")

    http.Handle("/favicon.ico", http.NotFoundHandler())
    http.HandleFunc("/", homepage)
    http.HandleFunc("/home", redirect1)
    http.HandleFunc("/index", redirect2)

    http.ListenAndServe(":8080", nil)
}
```

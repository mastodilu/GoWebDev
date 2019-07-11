# Cookie

I cookie sono piccoli file che contengono informazioni che il server può scrivere nella macchina del client se il client permette la scrittura di cookie.

Ogni volta che il client del browser invia una richiesta ad un certo server/dominio se il browser trova cookie per quel dominio li aggiunge alla request.

Il server riceve una request nel quale può trovare informazioni (ad esempio id univoci) associati a quel determinato utente.

## Workaround

Se i cookie non sono abilitati nel client, una delle possibili azioni è quella di specificare un parametro ID nell'url e usare https. Il server userà quel parametro per capire con quale utente sta comunicando.

NB: **Il nome di un cookie NON PUO' contenere spazi**.

## Read write cookies

### Script

```Go
package main

import (
    "fmt"
    "net/http"
)

func setMyCookie(w http.ResponseWriter, r *http.Request) {
    http.SetCookie(w, &http.Cookie{
        Name:  "supah-secret",
        Value: "Culo chi legge",
    })
    fmt.Fprintln(w, "Controlla i cookie.")
}

func main() {
    fmt.Println("Listening on port :8080")

    http.HandleFunc("/", setMyCookie)
    http.Handle("/favicon.ico", http.NotFoundHandler())

    http.ListenAndServe(":8080", nil)
}
```

### Output

![cookie nel client](img/001.png)

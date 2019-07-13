# Sessions

Le sessioni permettono di memorizzare informazioni associate a ID unici per poter associare determinate informazioni agli utenti.

## UUID Universally Unique ID

Sono ID standard nella programmazione web di 128 bit.

Il package consigliato per gestire questi ID e' [`github.com/satori/go.uuid`](https://godoc.org/github.com/satori/go.uuid#Size).

> `go get github.com/satori/go.uuid`

Un modo di usarli e' di salvarne il valore in un cookie, perche' i cookie sono interni al sito e non vengono inviati a terzi.

E' inoltre possibile settare quel cookie per usare solo https (`Secure=true`) e per essere accessibile solo tramite protocollo http (`HttpOnly=true`), niente javascript.

```Go
func home(w http.ResponseWriter, r *http.Request) {
    _, err := r.Cookie("uuid")
    if err != nil {
        // non esiste il cookie di sessione, lo creo
        id, err := uuid.NewV4()
        if err != nil {
            http.NotFoundHandler()
            log.Fatal(err)
        }
        session := http.Cookie{
            Name:  "uuid",
            Value: id.String(),
        }
        http.SetCookie(w, &session)
        fmt.Fprintln(w, "UUID created, check your cookies")
        return
    }

    fmt.Fprintln(w, "uuid already existent, check your cookies")
}
```

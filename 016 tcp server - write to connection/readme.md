# Readme

HTTP è un protocollo che sfrutta TCP.

TCP è un protocollo che controlla gli errori e comunica quando l'ascoltatore è pronto.

UDP invece comunica velocemente e senza controllare nulla.

## TCP server

Per creare un server HTTP, bisogna creare un server TCP che gestisce richieste formattate in un certo modo.

- [package net](https://golang.org/pkg/net/)
- [package net/http](https://golang.org/pkg/net/http/)

### Crea un TCP server

Per creare un TCP server si usa la funzione `net.Listen`.

```go
lnr, err := net.Listen("tcp", ":8080")
defer lnr.Close()
```

`lnr` è di tipo `*net.TCPListener` ed implementa questa interfaccia:

```Go
type Listener interface {
    // Accept waits for and returns the next connection to the listener.
    Accept() (Conn, error)

    // Close closes the listener.
    // Any blocked Accept operations will be unblocked and return errors.
    Close() error

    // Addr returns the listener's network address.
    Addr() Addr
}
```

Per gestire le richieste si usa `lnr.Accept(..)`

```Go
// gestisci le richieste
for {
    conn, err := lnr.Accept() // accetta le connessioni in arrivo
    if err != nil {
        panic(err)
    }

    handleRequest(conn)
}
```

Il metodo `Accept()` blocca il ciclo e resta in attesa di una connessione.

Per scrivere nella connection:

```Go
func handleRequest(conn net.Conn) {
    defer conn.Close()

    _, err := conn.Write([]byte("ciao dal server"))
    // handle error
    io.WriteString(conn, "\nCiao dal server\n")
    // handle error
    fmt.Fprintf(conn, "\n%s\n", "Ciao dal server")
    // handle error
}
```

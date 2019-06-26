# Readme

## Crea un server che resta in ascolto di nuove connessioni

```Go
    lnr, err := net.Listen("tcp", ":8080")
    // handle err
    defer lnr.Close()

    fmt.Println("Server started")

    for {
        conn, err := lnr.Accept()
        // handle err

        go func() {
            defer conn.Close()
            handleConnection(conn)
        }()
    }
}
```

### Scrivi al client

```Go
func handleConnection(conn net.Conn) {
    defer conn.Close()
    fmt.Fprintf(conn, "Hello from the server")
}
```

### Leggi dal client

```Go
func read(conn net.Conn) {
    fmt.Println("reading")
    reader := bufio.NewScanner(conn)
    var err error

    for reader.Scan() {
        if reader.Err() != nil {
            fmt.Println(err)
        }
        err := conn.SetDeadline(time.Now().Add(time.Second * 10))
        if err != nil {
            fmt.Println(err)
        }
        fmt.Printf("server: Client said %v\n", reader.Text())
    }
}
```

## Crea un client che si connette al server

```Go
conn, err := net.Dial("tcp", ":8080")
// handle err
defer conn.Close()
```

### Leggi dal server

```Go
func readFromServer(conn net.Conn) {
    reader := bufio.NewScanner(conn)
    for reader.Scan() {
        fmt.Println("Server said", reader.Text())
    }
}
```

### Scrivi al server

```Go
fmt.Fprintf(conn, "Hello from client")
```

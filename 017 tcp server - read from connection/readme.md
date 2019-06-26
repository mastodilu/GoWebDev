# Readme

## Leggi dalla connection

Per leggere dalla connection si può usare la funzione `bufio.Scan()`.

```Go
func read(conn net.Conn) string {
    fmt.Println("reading")
    reader := bufio.NewScanner(conn)
    var read string

    canRead := true
    for canRead {
        canRead = reader.Scan()
        if !canRead {
            err := reader.Err()
            switch err {
            case nil:
                // return read
                // EOF 💚 no error
            default:
                panic(err)
            }
        }
        read += reader.Text() + "\n"
    }
    fmt.Println("Stop listening.")
    return read
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    fmt.Println(read(conn))
    fmt.Printf("local addr %v remote addr %v\n", conn.LocalAddr(), conn.RemoteAddr())
}
```

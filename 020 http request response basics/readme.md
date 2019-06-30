# Readme

## HTTP request response foundations

### Leggere la request

```Go
func readFrom(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if scanner.Text() == "" {
			// headers are done
			// standard http: dopo gli header c'è una blank line
			break
		}
	}
}
```

E' necessario interrompere la lettura dopo l'**header**.
Stando allo standard http l'header termina con un `\n`.

### Fornire una response

```Go
func writeTo(conn net.Conn) {
	msg := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    <h1>CIAO DAL SERVER</h1>
</body>
</html>`
	fmt.Fprintln(conn, "HTTP/1.1 200 OK")
	fmt.Fprintf(conn, "Content-length: %d", len(msg))
	fmt.Fprintln(conn, "Content-Type: text/html")
	fmt.Fprintln(conn, "")
	fmt.Fprintf(conn, "%v\n\r", msg)
}
```

Una risposta valida deve fornire anche un header valido.

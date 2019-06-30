package main

import (
	"bufio"
	"fmt"
	"net"
)

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

func handleConnection(conn net.Conn) {
	defer conn.Close()
	readFrom(conn)
	writeTo(conn)
}

func main() {
	lsnr, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	fmt.Println("server started")
	for {
		conn, err := lsnr.Accept()
		if err != nil {
			panic(err)
		}

		go handleConnection(conn)
	}
}

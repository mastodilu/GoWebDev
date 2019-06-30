package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	var route string
	lineIndex := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if lineIndex == 0 {
			route = strings.Fields(scanner.Text())[1]
		}
		lineIndex++
		if scanner.Text() == "" {
			// headers are done
			// standard http: dopo gli header c'è una blank line
			break
		}
	}
	writeTo(conn, route)
}

func writeTo(conn net.Conn, route string) {
	msg := "<!DOCTYPE html>"
	msg += "<html lang=\"en\">"
	msg += "<head>"
	msg += "	<meta charset=\"UTF-8\">"
	msg += "	<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">"
	msg += "	<meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\">"
	msg += "	<title>Document</title>"
	msg += "</head>"
	msg += "<body>"
	msg += "<h1>" + route + "</h1>"
	msg += "</body>"
	msg += "</html>"
	fmt.Fprintln(conn, "HTTP/1.1 200 OK")
	fmt.Fprintf(conn, "Content-length: %d", len(msg))
	fmt.Fprintln(conn, "Content-Type: text/html")
	fmt.Fprintln(conn, "")
	fmt.Fprintf(conn, "%v\n\r", msg)
}

func handle(conn net.Conn) {
	defer conn.Close()
	handleConnection(conn)
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

		go handle(conn)
	}
}

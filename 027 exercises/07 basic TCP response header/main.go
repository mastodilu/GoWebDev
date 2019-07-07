package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func read(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		fmt.Println(scanner.Text())
	}
}

func serve(conn net.Conn) {

	body := "I see you connected"

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	// io.WriteString(conn, "\r\n")

	_, err := io.WriteString(conn, body)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Listening on port :8080 for TCP connections")
	lsr, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer lsr.Close()

	for {
		conn, err := lsr.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			defer conn.Close()
			read(conn)
			serve(conn)
		}()
	}
}

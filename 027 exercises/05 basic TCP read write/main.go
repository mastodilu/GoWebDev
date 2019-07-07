package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		fmt.Println(scanner.Text())
	}
	_, err := io.WriteString(conn, "I see you connected")
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
			handle(conn)
		}()
	}
}

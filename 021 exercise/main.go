package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handle(conn net.Conn) {
	// read
	scanner := bufio.NewScanner(conn)
	lineIndex := 0
	var host string

	for scanner.Scan() {
		if lineIndex == 1 {
			// 💥 url della request
			host = strings.Fields(scanner.Text())[1]
		}
		fmt.Printf("%v\n\r", scanner.Text())

		lineIndex++
		if scanner.Text() == "" {
			// fine header, esci dal ciclo
			break
		}
	}
	fmt.Printf("### host: %v\n\n", host)
}

func main() {
	lst, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer lst.Close()

	fmt.Println("Server started")
	for {
		conn, err := lst.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			defer conn.Close()
			handle(conn)
		}()
	}
}

package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func rot13(s string) string {
	rs := []rune(s)
	for i, r := range rs {
		if r >= 'a' && r <= 'm' {
			rs[i] = rs[i] + 13
		} else {
			rs[i] = rs[i] - 13
		}
	}
	return string(rs)
}

func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := strings.ToLower(scanner.Text())
		if text == "exit" {
			break
		} else {
			fmt.Printf("%v --> %v\n", text, rot13(text))
			fmt.Fprintf(conn, "%v --> %v\n", text, rot13(text))
		}
	}
}

func main() {
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			defer conn.Close()
			handleConnection(conn)
		}()
	}
}

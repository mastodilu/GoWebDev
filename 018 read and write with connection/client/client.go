package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// writeToServer(conn)
	readFromServer(conn)
}
func writeToServer(conn net.Conn) {
	fmt.Fprintf(conn, "Hello from client")
}

func readFromServer(conn net.Conn) {
	reader := bufio.NewScanner(conn)
	for reader.Scan() {
		fmt.Println("Server said", reader.Text())
	}
}

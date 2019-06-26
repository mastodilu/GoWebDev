package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

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

func write(conn net.Conn) {
	fmt.Fprintf(conn, "Ciao from server")
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// read(conn)
	write(conn)
}

func main() {
	lnr, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer lnr.Close()

	fmt.Println("Server started")

	for {
		conn, err := lnr.Accept()
		if err != nil {
			panic(err)
		}

		go func() {
			defer conn.Close()
			handleConnection(conn)
		}()
	}
}

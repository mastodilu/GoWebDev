package main

import (
	"fmt"
	"io"
	"net"
)

func handleRequest(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Handling a connection")

	_, err := conn.Write([]byte("ciao dal server"))
	if err != nil {
		panic(err)
	}

	_, err = io.WriteString(conn, "\nCiao dal server\n")
	if err != nil {
		panic(err)
	}

	_, err = fmt.Fprintf(conn, "\n%s\n", "Ciao dal server")
	if err != nil {
		panic(err)
	}
}

func main() {

	// crea il TCP listener per ascoltare le richieste in arrivo
	lnr, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer lnr.Close()
	fmt.Printf("tcp listener: type %T - %v - addr %v\n", lnr, lnr, lnr.Addr()) // *net.TCPListener - &{0xc00007e000}

	// gestisci le richieste
	for {
		conn, err := lnr.Accept() // accetta le connessioni in arrivo
		fmt.Printf("conn: type %T value %v\n", conn, conn)
		if err != nil {
			panic(err)
		}

		go handleRequest(conn)
	}

}

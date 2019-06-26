package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// handleConnection ascolta la connessione per delle keyword che permettono
// di interagire con il database in memory
func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		cmd := strings.Split(scanner.Text(), " ")
		cmdLen := len(cmd)
		if cmdLen < 1 {
			fmt.Printf("Command %v is invalid.\n\r", cmd)
		}
		switch cmd[0] {
		case "SET":
			if cmdLen < 3 {
				fmt.Printf("Command %v is invalid.\n\r", cmd)
			} else {
				memDB[cmd[1]] = cmd[2]
				fmt.Fprintf(conn, "set %v = %v\n\r", cmd[1], memDB[cmd[1]])
			}
		case "GET":
			if cmdLen < 2 {
				fmt.Printf("Command %v is invalid.\n\r", cmd)
			} else {
				value, ok := memDB[cmd[1]]
				if !ok {
					fmt.Fprintf(conn, "value %v is not in DB\n\r", cmd[1])
				} else {
					fmt.Fprintf(conn, "%v", value)
				}
			}
		case "DEL":
			if cmdLen < 2 {
				fmt.Printf("Command %v is invalid.\n\r", cmd)
			} else {
				if _, ok := memDB[cmd[1]]; !ok {
					fmt.Fprintf(conn, "value %v is not in DB\n\r", cmd[1])
				} else {
					delete(memDB, cmd[1])
					fmt.Fprintf(conn, "deleted key %v\n\r", cmd[1])
				}
			}

		case "PRINT":
			for k, v := range memDB {
				fmt.Fprintf(conn, "%v, %v\n\r", k, v)
			}
		default:
			fmt.Fprintf(conn, "Command %v is invalid.\n\r", cmd)
			fmt.Printf("Command %v is invalid.\n\r", scanner.Text())
		}
	}
}

func init() {
	memDB = make(map[string]string)
}

var memDB map[string]string

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server started")

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			defer conn.Close()
			handleConnection(conn)
		}()
	}

	fmt.Println("Server exited")
}

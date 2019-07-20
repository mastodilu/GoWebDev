package main

import (
	"database/sql"
	"fmt"
	"log"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	fmt.Println("Please enter DB password:")

	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	password := string(bytePassword)

	dbURL := fmt.Sprintf("root:%v@/myschema?charset=utf8", password) // OK
	// dbURL := fmt.Sprintf("root:%v@tcp(127.0.0.1:3306)/myschema?charset=utf8", password) 	// OK
	// dbURL := fmt.Sprintf("root:%v@tcp(localhost:3306)/myschema?charset=utf8", password) 	// OK
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("It's working")
}

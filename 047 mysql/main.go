package main

import (
	"database/sql"
	"fmt"
	"log"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh/terminal"
)

func create() {
	/*
	   CREATE TABLE Persons (
	       PersonID int,
	       LastName varchar(255),
	       FirstName varchar(255),
	       Address varchar(255),
	       City varchar(255)
	   );
	*/
	//query := `create table gotable`
}

func insert() {}

func read() {
	query := `select name from people;`
	fmt.Printf("query:\n\t%v\ncontent:\n", query)
	dbRows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	var rows []string
	for dbRows.Next() {
		var name string
		if err := dbRows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		rows = append(rows, name)

		fmt.Printf("\t%v\n", name)
	}
}

func readAllColumns() {
	query := `select * from people;`
	fmt.Printf("query:\n\t%v\ncontent:\n", query)
	dbRows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	var id, name string
	for dbRows.Next() {
		if err := dbRows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\t%v - %v\n", id, name)
	}
}

func update() {}

func delete() {}

func drop() {
	/*
	   DROP TABLE Shippers;
	*/
}

var db *sql.DB

func main() {
	fmt.Println("Please enter DB password:")
	var err error
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	password := string(bytePassword)

	dbURL := fmt.Sprintf("root:%v@/myschema?charset=utf8", password) // OK
	// dbURL := fmt.Sprintf("root:%v@tcp(127.0.0.1:3306)/myschema?charset=utf8", password) 	// OK
	// dbURL := fmt.Sprintf("root:%v@tcp(localhost:3306)/myschema?charset=utf8", password) 	// OK
	db, err = sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("you're connected to the DB.")

	readAllColumns()
}

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
	query := `CREATE TABLE DOG( 
		DogID INT(11) NOT NULL AUTO_INCREMENT,
		DogName VARCHAR(64),
		DogOwner VARCHAR(64),
		PRIMARY KEY (DogID)
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Done\n\t%v\n", query)
}

func insert1() {
	query := `INSERT INTO DOG (DogName, DogOwner) VALUES ('Fido', 'Mario');`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Done: %v\n", query)
}

func insert2() {
	dogNames := []string{"Fuffi", "Pluto", "Norberto"}
	dogOwners := []string{"Matteo", "Gennaro", "Hagrid"}
	for i := range dogNames {
		query := fmt.Sprintf("INSERT INTO DOG (DogName, DogOwner) VALUES ('%v', '%v');", dogNames[i], dogOwners[i])
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Done:\n\t%v\n", query)
	}
}

func insert3() {
	query := "INSERT INTO DOG (DogName, DogOwner) VALUES (?, ?);"
	dogNames := []string{"Fuffi", "Pluto", "Norberto"}
	dogOwners := []string{"Matteo", "Gennaro", "Hagrid"}
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var res sql.Result
	for i := range dogNames {
		res, err = stmt.Exec(dogNames[i], dogOwners[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	n, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted %v rows\n", n)
}

func read() {
	query := `select PersonName from PERSON;`
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
	query := `select * from PERSON;`
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

func update() {
	query := `UPDATE DOG SET DogName = 'Fidoh' WHERE DogName = 'Fido';`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Done: %v\n", query)
}

func delete() {
	query := `DELETE FROM DOG WHERE DogName = 'Fidoh';`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Done: %v\n", query)
}

func drop() {
	query := `DROP TABLE PERSON;`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Done: %v\n", query)
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

	// read()
	// readAllColumns()
	// create()
	// insert1()
	// insert2()
	// insert3()
	// update()
	// delete()
	// drop()

}

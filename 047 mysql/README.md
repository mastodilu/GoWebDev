# MySQL

Per usare un database MySQL gia' creato e' necessario importare i package:

- `database/sql`
- `github.com/go-sql-driver/mysql`

che va prima installato con `go get -u github.com/go-sql-driver/mysql`.

```Go
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

db, err := sql.Open("mysql", "user:password@/dbname")
```

Il package del driver viene importato con l'**alias vuoto `_`** perche' viene usato solo in fase di configurazione e mai all'interno del codice.

`Open` non apre la connessione, quindi per verificare che tutto sia andato a buon fine si usa:

```Go
func (c *Conn) PingContext(ctx context.Context) error
```

## Esempio

```Go
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
```

Il package `"golang.org/x/crypto/ssh/terminal"` permette di far scrivere la password all'utente senza che venga stampata a video durante la digitazione.

NB: il nome del database da usare e' quello che in **MySQL workbench** si chiama `schema`:

![schema database name](schema.png)

## Read

```Go
func read() {
    query := `select name from people;`
    dbRows, err := db.Query(query)
    if err != nil {
        log.Fatal(err)
    }

    for dbRows.Next() {
        var name string
        if err := dbRows.Scan(&name); err != nil {
            log.Fatal(err)
        }
        fmt.Println(name)
    }
}
```

Si esegue la `SELECT` cosi':

```Go
query := `select name from people;`
dbRows, err := db.Query(query)
/* check error */
```

Il metodo `(db *DB) Query` vuole come parametro la query e restituisce un puntatore a `Rows`:

```Go
func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
```

Per scorrere le righe del risultato si usa

```Go
func (rs *Rows) Next() bool
```

che prepara la lettura della riga corrente con il metodo

```Go
func (rs *Rows) Scan(dest ...interface{}) error
```

`Scan` vuole come parametri dei puntatori, uno per colonna letta dalla tabella. Ad esempio:

```GO
query := `select * from people;`
dbRows, err := db.Query(query)
// error ...
var id, name string
for dbRows.Next() {
    dbRows.Scan(&id, &name) // <--
    // error...
}
```

## Create

```Go
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
```

## Insert

```Go
func insert1() {
    query := `INSERT INTO DOG (DogName, DogOwner) VALUES ('Fido', 'Mario');`
    _, err := db.Exec(query)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Done: %v\n", query)
}
```

## Eseguire molte query

Ci sono due modi efficienti:

1. usare la funzione `fmt.Sprintf`
2. usare i prepared statement

### `fmt.SPrintf`

```Go
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
```

### Prepared statement

```Go
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
```

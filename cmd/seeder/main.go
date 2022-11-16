package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	os.Remove("product.db")

	db, err := sql.Open("sqlite3", "product.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	CREATE TABLE products (ID INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, price REAL);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

}

func seeder(db *sql.DB) {

}

package main

import (
	"database/sql"
	"fmt"
	"log"

	// "github.com/davecgh/go-spew/spew"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "zip.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, code, city, state FROM places")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var code, city, state string
		err = rows.Scan(&id, &code, &city, &state)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, code, city, state)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

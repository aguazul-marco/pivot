package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "zip.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := seed(db); err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, code, city, state FROM places")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	type place struct {
		id    int
		code  string
		city  string
		state string
	}

	var places []place

	for rows.Next() {
		var p place
		err = rows.Scan(&p.id, &p.code, &p.city, &p.state)
		if err != nil {
			log.Fatal(err)
		}
		places = append(places, p)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(places)
}

func seed(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("INSERT INTO places(code, city, state) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for i := 0; i < 5; i++ {
		_, err := stmt.Exec(fmt.Sprintf("code-%d", i), fmt.Sprintf("city-%d", i), fmt.Sprintf("state-%d", i))
		if err != nil {
			return err
		}
		// spew.Dump(r)
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

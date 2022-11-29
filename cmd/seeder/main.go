package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var Products []Product

func initProduct() {
	data, err := os.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(data, &Products); err != nil {
		log.Fatal(err)
	}
}

func main() {
	initProduct()
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

	if err := seed(db); err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, name, price FROM products LIMIT 5")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, p)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	Products = products

	for _, p := range Products {
		fmt.Printf("ID: %v\nName: %v\nPrice: %v\n", p.ID, p.Name, p.Price)
	}

}

func seed(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("INSERT INTO products (name, price) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for i := 0; i < 5; {
		for _, product := range Products {
			_, err = stmt.Exec(product.Name, product.Price)
			if err != nil {
				return err
			}
			i++
		}

	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

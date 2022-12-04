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

func initProducts(path string) ([]Product, error) {
	var products []Product
	data, err := os.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(data, &products); err != nil {
		return []Product{}, err
	}

	return products, nil
}

func main() {
	products, err := initProducts("products.json")
	if err != nil {
		log.Fatal(err)
	}
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

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("INSERT INTO products (name, price) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, product := range products {
		_, err = stmt.Exec(product.Name, product.Price)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, name, price FROM products LIMIT 5")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			fmt.Print(err)
			return
		}
		fmt.Println(p.ID, p.Name, p.Price)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}

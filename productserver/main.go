package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var db *sql.DB

func InitProduct(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Printf("Could not connect with database: %v", err)
	}

	fmt.Printf("Connected with database successfully\n")

	return db
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	lt := query.Get("limit")

	limit, err := strconv.Atoi(lt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	if limit == 0 {
		limit = 200
	}

	rows, err := db.Query("SELECT id, name, price FROM products ORDER BY id ASC LIMIT ?", limit)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	if err := json.NewEncoder(w).Encode(products); err != nil {
		log.Printf("error occured while encoding: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func GetProductID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Not a numeric value: %v", err)
	}

	row := db.QueryRow("SELECT id, name, price FROM products WHERE id = ?", id)
	var p Product
	switch err := row.Scan(&p.ID, &p.Name, &p.Price); err {
	case sql.ErrNoRows:
		w.WriteHeader(http.StatusNotFound)
		log.Println(err)
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	if p.ID == 0 && p.Name == "" {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("No product exists with id %v", id)
		return
	}

	if err := json.NewEncoder(w).Encode(p); err != nil {
		log.Printf("error encoding product: %v", id)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("OK")
	w.WriteHeader(http.StatusOK)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	tx, err := db.Begin()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}
	stmt, err := tx.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}
	defer stmt.Close()

	results, err := stmt.Exec(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	rowsAffected, err := results.RowsAffected()
	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("No product exists with id %v", id)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	fmt.Println("OK")
	w.WriteHeader(http.StatusOK)
}

func AddNewProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var p Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Printf("error occured while decoding: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if p.Name == "" || p.Price == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}
	stmt, err := tx.Prepare("INSERT INTO products (name, price) VALUES (?,?)")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Name, p.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	fmt.Println("Created")
	w.WriteHeader(http.StatusCreated)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Not a numeric value: %v", err)
	}

	var p Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Printf("error occured while decoding: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if p.Name == "" || p.Price == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}
	stmt, err := tx.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(p.Name, p.Price, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("No product exists with id %v", id)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	fmt.Println("Updated")
	w.WriteHeader(http.StatusOK)
}

func main() {
	var dbFile string
	flag.StringVar(&dbFile, "db", "product.db", "find products.db")
	flag.Parse()

	db = InitProduct(dbFile)

	r := mux.NewRouter()
	r.HandleFunc("/products", AddNewProduct).Methods("POST")
	r.HandleFunc("/products/{id}", UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", DeleteProduct).Methods("DELETE")
	r.HandleFunc("/products/{id}", GetProductID).Methods("GET")
	r.HandleFunc("/products", GetProducts).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

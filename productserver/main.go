package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

var Products []Product

func InitProduct() {
	data, err := os.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(data, &Products); err != nil {
		log.Fatal(err)
	}
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(Products); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func GetProductID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]

	for _, p := range Products {
		if strconv.Itoa(p.ID) == key {
			if err := json.NewEncoder(w).Encode(p); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Println("Product not found error:")
			fmt.Println(http.StatusNotFound)
		}
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	for index, p := range Products {
		if strconv.Itoa(p.ID) == id {
			Products = append(Products[:index], Products[index+1:]...)
			break
		}
	}
	if err := json.NewEncoder(w).Encode(Products); err != nil {
		log.Fatal(err)
	}
}

func AddNewProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ID = len(Products) + 2
	Products = append(Products, product)
	if err := json.NewEncoder(w).Encode(product); err != nil {
		log.Fatal(err)
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	for i, p := range Products {
		if strconv.Itoa(p.ID) == id {
			Products = append(Products[:i], Products[i+1:]...)
			var product Product
			_ = json.NewDecoder(r.Body).Decode(&product)
			ID, err := strconv.Atoi(id)
			if err != nil {
				log.Fatal(err)
			}
			product.ID = ID
			Products = append(Products, product)
			if err = json.NewEncoder(w).Encode(product); err != nil {
				log.Fatal(err)
			}
			return
		}
	}
}

func main() {

	InitProduct()

	r := mux.NewRouter()
	r.HandleFunc("/products", GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", GetProductID).Methods("GET")
	r.HandleFunc("/products/{id}", DeleteProduct).Methods("DELETE")
	r.HandleFunc("/products/{id}", AddNewProduct).Methods("POST")
	r.HandleFunc("/products/{id}", UpdateProduct).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", r))
}

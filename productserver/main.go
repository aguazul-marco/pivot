package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
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
		log.Printf("error occured while encoding: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func GetProductID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("error converting id to int: %v", err)
	}
	p := productFinder(id)
	if p == nil {
		log.Printf("no product exists with id %v", id)
		return
	}
	if err := json.NewEncoder(w).Encode(p); err != nil {
		log.Printf("error encoding product: %v", id)
		w.WriteHeader(http.StatusInternalServerError)
		return
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
		log.Printf("error occured while encoding: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func AddNewProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		log.Printf("error occured while decoding: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	product.ID = len(Products) + 2
	Products = append(Products, product)
	if err := json.NewEncoder(w).Encode(product); err != nil {
		log.Printf("error occured while encoding: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
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
			if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
				log.Printf("error occured while decoding: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
			}

			i, err := strconv.Atoi(id)
			if err != nil {
				log.Printf("error occured while converting int: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
			}

			product.ID = i
			Products = append(Products, product)
			if err = json.NewEncoder(w).Encode(product); err != nil {
				log.Printf("error occured while encoding: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
	}
}

func productFinder(id int) *Product {
	for _, p := range Products {
		if p.ID == id {
			return &p
		}
	}
	return nil
}

func main() {

	InitProduct()

	r := mux.NewRouter()
	r.HandleFunc("/products/{id}", AddNewProduct).Methods("POST")
	r.HandleFunc("/products/{id}", UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", DeleteProduct).Methods("DELETE")
	r.HandleFunc("/products/{id}", GetProductID).Methods("GET")
	r.HandleFunc("/products", GetProducts).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

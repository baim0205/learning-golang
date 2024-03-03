package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Product struct {
	ID    int
	Name  string
	Price float64
}

var products []Product

func CreateProduct(product Product) {
	products = append(products, product)
}

func ReadProducts() []Product {
	return products
}

func UpdateProduct(id int, updatedProduct Product) {
	for i, product := range products {
		if product.ID == id {
			products[i] = updatedProduct
			return
		}
	}
}

func DeleteProduct(id int) {
	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			return
		}
	}
}

func main() {
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productsJSON, _ := json.Marshal(ReadProducts())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(productsJSON)

		case http.MethodPost:
			var newProduct Product
			err := json.NewDecoder(r.Body).Decode(&newProduct)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			CreateProduct(newProduct)
			w.WriteHeader(http.StatusCreated)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		idStr := r.URL.Path[len("/products/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for _, product := range ReadProducts() {
			if product.ID == id {
				productJSON, _ := json.Marshal(product)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(productJSON)
				return
			}
		}

		http.NotFound(w, r)
	})

	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

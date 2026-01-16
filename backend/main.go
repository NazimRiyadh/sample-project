package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var products []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

func Handle_cors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
}

func TypeCheck(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		http.Error(w, "Make a "+r.Method+" request", 400)
		return false
	}
	return true
}

func PreflightRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
	}
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	Handle_cors(w, r)
	PreflightRequest(w, r)
	if TypeCheck(w, r, "GET") == false {
		return
	}
	json.NewEncoder(w).Encode(products)
}

func addProducts(w http.ResponseWriter, r *http.Request) {
	Handle_cors(w, r)
	PreflightRequest(w, r)

	if TypeCheck(w, r, "GET") == false {
		return
	}

	var new_product Product
	json.NewDecoder(r.Body).Decode(&new_product)

	new_product.ID = len(products) + 1
	products = append(products, new_product)

	json.NewEncoder(w).Encode(products)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/products", getProducts)
	mux.HandleFunc("/addProducts", addProducts)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	} else {
		fmt.Println("Server started on port 8080")
	}
}

func init() {
	prod1 := Product{ID: 1, Title: "orange", Description: "my father's fav fruit", Price: 129.99, ImgURL: "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS"}
	prod2 := Product{ID: 2, Title: "apple", Description: "tasty to eat while fever", Price: 199.99, ImgURL: "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528"}
	prod3 := Product{ID: 3, Title: "Watermelon", Description: "Tasty & Juicy", Price: 89.50, ImgURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s"}

	products = append(products, prod1, prod2, prod3)
}

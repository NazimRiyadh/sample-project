package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Parse ID from URL path (e.g., /products/1)
	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[len(parts)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Decode the product from request body
	var product database.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Update the product
	database.Update(id, product)

	util.SendData(w, "Product updated successfully", http.StatusOK)
}

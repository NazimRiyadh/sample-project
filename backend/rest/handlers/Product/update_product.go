package Product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type ReqUpdateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Parse ID from URL path (e.g., /products/1)
	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[len(parts)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Decode the product from request body
	var product ReqUpdateProduct
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Update the product
	updatedProduct, err := h.productRepo.Update(repo.Product{
		ID:          id,
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		ImgURL:      product.ImgURL,
	})
	if err != nil {
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		return
	}
	util.SendData(w, updatedProduct, http.StatusOK)
}

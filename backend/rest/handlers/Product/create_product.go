package Product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type ReqCreateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

func (h *ProductHandler) AddProducts(w http.ResponseWriter, r *http.Request) {

	var new_product ReqCreateProduct
	err := json.NewDecoder(r.Body).Decode(&new_product)
	if err != nil {
		util.SendData(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	createdProduct, err := h.productRepo.Create(repo.Product{
		Title:       new_product.Title,
		Description: new_product.Description,
		Price:       new_product.Price,
		ImgURL:      new_product.ImgURL,
	})
	if err != nil {
		util.SendData(w, "Error creating product", http.StatusInternalServerError)
		return
	}
	util.SendData(w, createdProduct, http.StatusOK)
}

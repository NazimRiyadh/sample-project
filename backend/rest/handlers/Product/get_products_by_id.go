package Product

import (
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
)

func (h *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	Id := r.PathValue("productId")

	id, err := strconv.Atoi(Id)
	if err != nil {
		fmt.Println("Error converting productId to int")
		return
	}
	data, err := h.productRepo.Get(id)
	if data != nil {
		util.SendData(w, data, http.StatusOK)
		return
	}
	if err != nil {
		http.Error(w, "Error fetching product", http.StatusInternalServerError)
		return
	}
	util.SendData(w, "Product not found", http.StatusNotFound)
}

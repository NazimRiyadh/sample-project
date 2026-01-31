package Product

import (
	"ecommerce/util"
	"net/http"
)

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productRepo.List()
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}
	util.SendData(w, products, http.StatusOK)
}

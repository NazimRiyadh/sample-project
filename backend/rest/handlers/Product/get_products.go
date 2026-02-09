package Product

import (
	"ecommerce/util"
	"fmt"
	"net/http"
)

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productRepo.List()
	if err != nil {
		fmt.Println("DB Error in GetProducts:", err)
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}
	util.SendData(w, products, http.StatusOK)
}

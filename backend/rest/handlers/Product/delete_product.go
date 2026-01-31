package Product

import (
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
)

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	Id := r.PathValue("productId")

	id, err := strconv.Atoi(Id)
	if err != nil {
		fmt.Println("Error converting productId to int")
		return
	}
	err = h.productRepo.Delete(id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, "Product deleted successfully", http.StatusOK)
}

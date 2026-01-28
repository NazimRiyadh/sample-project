package Product

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
)

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products := database.List()
	util.SendData(w, products, http.StatusOK)
}

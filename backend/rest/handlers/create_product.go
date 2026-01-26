package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

func AddProducts(w http.ResponseWriter, r *http.Request) {
	var new_product database.Product
	err := json.NewDecoder(r.Body).Decode(&new_product)
	if err != nil {
		util.SendData(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	database.Store(new_product)
	util.SendData(w, database.List(), http.StatusOK)
}

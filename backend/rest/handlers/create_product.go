package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

func AddProducts(w http.ResponseWriter, r *http.Request) {
	var new_product database.Product
	json.NewDecoder(r.Body).Decode(&new_product)

	database.Store(new_product)
	util.SendData(w, database.List(), http.StatusOK)
}

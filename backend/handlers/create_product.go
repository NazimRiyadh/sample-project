package handlers

import (
	"ecommerce/product"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

func AddProducts(w http.ResponseWriter, r *http.Request) {
	var new_product product.Product
	json.NewDecoder(r.Body).Decode(&new_product)

	new_product.ID = len(product.Products) + 1
	product.Products = append(product.Products, new_product)
	util.SendData(w, product.Products, r)
}

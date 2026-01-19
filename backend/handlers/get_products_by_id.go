package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")

	id, err := strconv.Atoi(productId)
	if err != nil {
		fmt.Println("Error converting productId to int")
		return
	}

	for _, product := range database.Products {
		if product.ID == id {
			util.SendData(w, product, http.StatusOK)
			return
		}
	}
	util.SendData(w, nil, http.StatusNotFound)
}

package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	Id := r.PathValue("productId")

	id, err := strconv.Atoi(Id)
	if err != nil {
		fmt.Println("Error converting productId to int")
		return
	}
	data := database.Get(id)
	if data != nil {
		util.SendData(w, data, http.StatusOK)
		return
	}
	util.SendData(w, "Product not found", http.StatusNotFound)
}

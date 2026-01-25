package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	Id := r.PathValue("productId")

	id, err := strconv.Atoi(Id)
	if err != nil {
		fmt.Println("Error converting productId to int")
		return
	}
	database.Delete(id)

	util.SendData(w, "Product deleted successfully", http.StatusOK)
}

package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var new_user database.Users
	err := json.NewDecoder(r.Body).Decode(&new_user)
	if err != nil {
		util.SendData(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	database.StoreUser(new_user)
	util.SendData(w, "User Created Successfully", http.StatusOK)
}

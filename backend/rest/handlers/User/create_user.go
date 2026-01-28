package User

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var new_user database.Users

	err := json.NewDecoder(r.Body).Decode(&new_user)
	if err != nil {
		util.SendData(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	database.StoreUser(new_user)
	// config := config.GetConfig()
	// util.CreateJWT(util.Payload{
	// 	ID:        new_user.ID,
	// 	Firstname: new_user.Firstname,
	// 	Lastname:  new_user.Lastname,
	// 	Email:     new_user.Email,
	// 	IsOwner:   new_user.IsOwner}, config.JWT_SECRET)
	util.SendData(w, "User created successfully", http.StatusOK)
}

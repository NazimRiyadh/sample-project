package User

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type ReqCreateUser struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsOwner   bool   `json:"is_owner"`
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var new_user ReqCreateUser

	err := json.NewDecoder(r.Body).Decode(&new_user)
	if err != nil {
		util.SendData(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	_, err = h.userRepo.Create(repo.User{
		Firstname: new_user.Firstname,
		Lastname:  new_user.Lastname,
		Email:     new_user.Email,
		Password:  new_user.Password,
		IsOwner:   new_user.IsOwner,
	})
	if err != nil {
		util.SendData(w, "Error creating user", http.StatusInternalServerError)
		return
	}
	// config := config.GetConfig()
	// util.CreateJWT(util.Payload{
	// 	ID:        new_user.ID,
	// 	Firstname: new_user.Firstname,
	// 	Lastname:  new_user.Lastname,
	// 	Email:     new_user.Email,
	// 	IsOwner:   new_user.IsOwner}, config.JWT_SECRET)
	util.SendData(w, "User created successfully", http.StatusOK)
}

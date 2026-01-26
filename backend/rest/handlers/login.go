package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type ReqLogin struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var login ReqLogin
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		util.SendData(w, "Invalid json to decode", http.StatusBadRequest)
	}
	user := database.GetUser(login.Id)
	if user == nil {
		util.SendData(w, "Invalid username or password", http.StatusUnauthorized)
	} else {
		if user.Password == login.Password {
			util.SendData(w, "Login successful", http.StatusOK)
		} else {
			util.SendData(w, "Invalid username or password", http.StatusUnauthorized)
		}
	}

}

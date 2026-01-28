package User

import (
	"ecommerce/config"
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type ReqLogin struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
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
			config := config.GetConfig()
			jwt, err := util.CreateJWT(util.Payload{
				ID:        user.ID,
				Firstname: user.Firstname,
				Lastname:  user.Lastname,
				Email:     user.Email,
				IsOwner:   user.IsOwner,
			}, config.JWT_SECRET)
			if err != nil {
				util.SendData(w, "Error creating token", http.StatusInternalServerError)
				return
			}
			util.SendData(w, jwt, http.StatusOK)
		} else {
			util.SendData(w, "Invalid username or password", http.StatusUnauthorized)
		}
	}

}

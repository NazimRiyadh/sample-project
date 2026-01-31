package User

import (
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var login ReqLogin
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		util.SendData(w, "Invalid json to decode", http.StatusBadRequest)
	}

	user, err := h.userRepo.Find(login.Email, login.Password)
	if err != nil {
		util.SendData(w, "Invalid username or password", http.StatusUnauthorized)
	} else {
		jwt, err := util.CreateJWT(util.Payload{
			ID:        user.ID,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Email:     user.Email,
			IsOwner:   user.IsOwner,
		}, h.config.JWT_SECRET)
		if err != nil {
			util.SendData(w, "Error creating token", http.StatusInternalServerError)
			return
		}
		util.SendData(w, jwt, http.StatusOK)
	}
}

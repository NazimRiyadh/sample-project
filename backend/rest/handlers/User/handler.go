package User

import (
	"ecommerce/config"
	"ecommerce/repo"
)

type UserHandler struct {
	userRepo repo.UserRepo
	config   *config.Config
}

func NewUserHandler(userRepo repo.UserRepo, config *config.Config) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
		config:   config,
	}
}

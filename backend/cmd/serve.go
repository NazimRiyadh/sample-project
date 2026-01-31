package cmd

import (
	"ecommerce/config"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/Product"
	"ecommerce/rest/handlers/Test"
	"ecommerce/rest/handlers/User"
	"ecommerce/rest/middlewares"
)

func Serve() {
	config := config.GetConfig()

	userRepo := repo.NewUserRepo()
	productRepo := repo.NewProductRepo()

	middlewareConfig := middlewares.NewMiddleware(config)

	producthandler := Product.NewProductHandler(middlewareConfig, productRepo)
	userHandler := User.NewUserHandler(userRepo, config)
	testHandler := Test.NewTestHandler()

	server := rest.NewServer(producthandler, userHandler, testHandler, config)
	server.Server()
}

package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/Product"
	"ecommerce/rest/handlers/Test"
	"ecommerce/rest/handlers/User"
	"ecommerce/rest/middlewares"
)

func Serve() {
	config := config.GetConfig()
	middlewareConfig := middlewares.NewMiddleware(config)
	producthandler := Product.NewProductHandler(middlewareConfig)
	userHandler := User.NewUserHandler()
	testHandler := Test.NewTestHandler()
	server := rest.NewServer(producthandler, userHandler, testHandler, config)
	server.Server()
}

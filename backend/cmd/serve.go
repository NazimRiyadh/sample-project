package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/Product"
	"ecommerce/rest/handlers/Test"
	"ecommerce/rest/handlers/User"
)

func Serve() {
	config := config.GetConfig()

	producthandler := Product.NewProductHandler()
	userHandler := User.NewUserHandler()
	testHandler := Test.NewTestHandler()
	server := rest.NewServer(producthandler, userHandler, testHandler, &config)
	server.Server()
}

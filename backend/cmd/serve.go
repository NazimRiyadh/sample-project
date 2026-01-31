package cmd

import (
	"ecommerce/Infra/db"
	"ecommerce/config"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/Product"
	"ecommerce/rest/handlers/Test"
	"ecommerce/rest/handlers/User"
	"ecommerce/rest/middlewares"
	"log"
)

func Serve() {
	config := config.GetConfig()

	dbCon, error := db.NewConnection()
	if error != nil {
		log.Fatal("Failed to connect to database:", error)
	}

	userRepo := repo.NewUserRepo()
	productRepo := repo.NewProductRepo()

	middlewareConfig := middlewares.NewMiddleware(config)

	producthandler := Product.NewProductHandler(middlewareConfig, productRepo)
	userHandler := User.NewUserHandler(userRepo, config)
	testHandler := Test.NewTestHandler()

	server := rest.NewServer(producthandler, userHandler, testHandler, config)
	server.Server()
}

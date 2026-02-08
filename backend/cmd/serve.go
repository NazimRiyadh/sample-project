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
	"fmt"
	"os"
)

func Serve() {
	config := config.GetConfig()

	db, error := db.NewConnection()
	if error != nil {
		fmt.Println("Error connecting to database:", error)
		os.Exit(1)
	}

	userRepo := repo.NewUserRepo(db)
	productRepo := repo.NewProductRepo(db)

	middlewareConfig := middlewares.NewMiddleware(config)

	producthandler := Product.NewProductHandler(middlewareConfig, productRepo)
	userHandler := User.NewUserHandler(userRepo, config)
	testHandler := Test.NewTestHandler()

	server := rest.NewServer(producthandler, userHandler, testHandler, config)
	server.Server()
}

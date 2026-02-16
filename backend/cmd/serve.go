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

	dbConn, err := db.NewConnection(&config.DBConfig)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		os.Exit(1)
	}

	err = db.Migrate(dbConn, "./migrations")
	if err != nil {
		fmt.Println("Error migrating database:", err)
		os.Exit(1)
	}

	userRepo := repo.NewUserRepo(dbConn)
	productRepo := repo.NewProductRepo(dbConn)

	middlewareConfig := middlewares.NewMiddleware(config)

	producthandler := Product.NewProductHandler(middlewareConfig, productRepo)
	userHandler := User.NewUserHandler(userRepo, config)
	testHandler := Test.NewTestHandler()

	server := rest.NewServer(producthandler, userHandler, testHandler, config)
	server.Server()
}

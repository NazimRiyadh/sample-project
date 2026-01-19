package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middlewares"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()
	manager := middlewares.NewManager()
	manager.Use(middlewares.Logger, middlewares.Hudai, middlewares.Arekta)

	mux.Handle("GET /test", manager.With(http.HandlerFunc(handlers.Test)))
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.AddProducts)))
	mux.Handle("GET /products/{productId}", manager.With(http.HandlerFunc(handlers.GetProductById)))

	globalRouter := global_router.GlobalRouter(mux)
	err := http.ListenAndServe(":8080", globalRouter)
	fmt.Println("Server started running on port 8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

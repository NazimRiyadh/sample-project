package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()
	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", (http.HandlerFunc(handlers.AddProducts)))

	globalRouter := global_router.GlobalRouter(mux)
	err := http.ListenAndServe(":8080", globalRouter)
	fmt.Println("Server started running on port 8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

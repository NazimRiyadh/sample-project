package rest

import (
	"ecommerce/rest/handlers"
	"ecommerce/rest/middlewares"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager middlewares.Manager) {
	mux.Handle("GET /test", manager.With(http.HandlerFunc(handlers.Test)))
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.AddProducts)))
	mux.Handle("GET /products/{productId}", manager.With(http.HandlerFunc(handlers.GetProductById)))
	mux.Handle("PUT /products/{productId}", manager.With(http.HandlerFunc(handlers.UpdateProduct)))
	mux.Handle("DELETE /products/{productId}", manager.With(http.HandlerFunc(handlers.DeleteProduct)))

	mux.Handle("POST /users", manager.With(http.HandlerFunc(handlers.CreateUser)))
	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(handlers.Login)))
}

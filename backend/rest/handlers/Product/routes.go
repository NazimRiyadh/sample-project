package Product

import (
	"ecommerce/rest/middlewares"
	"net/http"
)

func (h *ProductHandler) RegProductRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(h.GetProducts)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(h.AddProducts), h.middlwareConfig.Authenticate_jwt))
	mux.Handle("GET /products/{productId}", manager.With(http.HandlerFunc(h.GetProductById)))
	mux.Handle("PUT /products/{productId}", manager.With(http.HandlerFunc(h.UpdateProduct), h.middlwareConfig.Authenticate_jwt))
	mux.Handle("DELETE /products/{productId}", manager.With(http.HandlerFunc(h.DeleteProduct), h.middlwareConfig.Authenticate_jwt))
}

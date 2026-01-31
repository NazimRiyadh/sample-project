package Product

import (
	"ecommerce/repo"
	"ecommerce/rest/middlewares"
)

type ProductHandler struct {
	middlwareConfig *middlewares.MiddlewareConfig
	productRepo     repo.ProductRepo
}

func NewProductHandler(middlewareConfig *middlewares.MiddlewareConfig,
	productRepo repo.ProductRepo) *ProductHandler {
	return &ProductHandler{
		middlwareConfig: middlewareConfig,
		productRepo:     productRepo,
	}
}

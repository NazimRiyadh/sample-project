package Product

import "ecommerce/rest/middlewares"

type ProductHandler struct {
	middlwareConfig *middlewares.MiddlewareConfig
}

func NewProductHandler(middlewareConfig *middlewares.MiddlewareConfig) *ProductHandler {
	return &ProductHandler{
		middlwareConfig: middlewareConfig,
	}
}

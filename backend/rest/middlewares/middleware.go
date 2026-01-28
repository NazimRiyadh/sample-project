package middlewares

import "ecommerce/config"

type MiddlewareConfig struct {
	cnf *config.Config
}

func NewMiddleware(cnf *config.Config) *MiddlewareConfig {
	return &MiddlewareConfig{
		cnf: cnf,
	}
}

package Test

import (
	"ecommerce/rest/middlewares"
	"net/http"
)

func (h *TestHandler) RegtestHandler(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("GET /test", manager.With(http.HandlerFunc(h.Test)))
}

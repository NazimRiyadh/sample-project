package User

import (
	"ecommerce/rest/middlewares"
	"net/http"
)

func (h *UserHandler) RegUserRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("POST /users", manager.With(http.HandlerFunc(h.CreateUser)))
	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(h.Login)))
}

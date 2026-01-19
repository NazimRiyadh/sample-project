package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	Middlewares []Middleware
}

//will help to provide a object of manager
func NewManager() *Manager {
	mngr := Manager{
		Middlewares: []Middleware{},
	}
	return &mngr
}

func (m *Manager) With(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		handler := next

		for i := len(middlewares) - 1; i >= 0; i-- {
			middleware := middlewares[i]
			handler = middleware(handler)
		}
		return handler
	}
}

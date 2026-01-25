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

//for appending the middleares
func (m *Manager) Use(middlewares ...Middleware) {
	m.Middlewares = append(m.Middlewares, middlewares...)
}

func (m *Manager) With(handlers http.Handler, middlewares ...Middleware) http.Handler {

	// for _, middlewares := range middlewares {
	// 	handlers = middlewares(handlers)
	// }

	for _, middleware := range m.Middlewares {
		handlers = middleware(handlers)
	}
	return handlers
}

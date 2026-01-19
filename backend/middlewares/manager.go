package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	Middlewares []Middleware
}

func NewManager() *Manager {
	mngr := Manager{
		Middlewares: []Middleware{},
	}
	return &mngr
}

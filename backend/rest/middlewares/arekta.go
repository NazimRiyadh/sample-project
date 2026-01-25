package middlewares

import (
	"fmt"
	"net/http"
)

func Arekta(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Ami Arekta Middleware")
		next.ServeHTTP(w, r)
	})

}

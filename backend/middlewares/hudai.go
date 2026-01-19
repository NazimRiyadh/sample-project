package middlewares

import (
	"fmt"
	"net/http"
)

func Hudai(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Ami Hudai aisi vai, cholen shuru kori!")
		next.ServeHTTP(w, r)
	})
}

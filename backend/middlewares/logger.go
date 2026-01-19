package middlewares

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Println("Ami Logger middleware start kortesi")
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		log.Printf(r.Method, r.RequestURI, duration)
		log.Println("Ami Logger middleware sesh kortesi")
	})

}

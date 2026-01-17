package global_router

import "net/http"

func GlobalRouter(mux *http.ServeMux) http.Handler {
	HandleCors := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		} else {
			mux.ServeHTTP(w, r)
		}
	}

	return http.HandlerFunc(HandleCors)
}

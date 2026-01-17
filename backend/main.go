package main

import "ecommerce/cmd"

// Removed - CORS handling is now done in GlobalRoute middleware

// func TypeCheck(w http.ResponseWriter, r *http.Request, method string) bool {
// 	if r.Method != method {
// 		http.Error(w, "Make a "+r.Method+" request", 400)
// 		return false
// 	}
// 	return true
// }

// func PreflightRequest(w http.ResponseWriter, r *http.Request) {
// 		if r.Method == "OPTIONS" {
// 		w.WriteHeader(200)
// 		return
// 	}
// }

func main() {
	cmd.Serve()
}

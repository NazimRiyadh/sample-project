package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/config"
	"ecommerce/util"
	"fmt"
	"net/http"
	"strings"
)

func Authenticate_jwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//parse jwt
		//split the jwt use space
		//take the header and the payload
		//hash=hmac-sha256(header, payload, secret from my env)
		//take the signature part from jwt
		//if hash==signature -> go forward else unauthorized

		fmt.Println("-------------AUTHENTICAION-----------------")

		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Header not provided", http.StatusUnauthorized)
			return
		}

		header_split := strings.Split(header, " ")

		if len(header_split) != 2 {
			http.Error(w, "Provide valid header", http.StatusUnauthorized)
			return
		}

		jwt := header_split[1]
		jwt_split := strings.Split(jwt, ".")
		if len(jwt_split) != 3 {
			http.Error(w, "Not a JWT token", http.StatusUnauthorized)
		}
		jwt_header := jwt_split[0]
		jwt_payload := jwt_split[1]
		jwt_signature := jwt_split[2]

		//concat payload and header = message
		message := jwt_header + "." + jwt_payload

		//message to bytearray
		byte_arr_message := []byte(message)

		config := config.GetConfig()

		//secret to bytesecret
		byte_arr_secret := []byte(config.JWT_SECRET)

		h := hmac.New(sha256.New, byte_arr_secret)
		h.Write(byte_arr_message)

		signature := h.Sum(nil)

		base64_signature := util.Base64encoding(signature)

		if base64_signature != jwt_signature {
			http.Error(w, "halarpo, tui hacker", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	ID        int    `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
	IsOwner   bool   `json:"is_owner"`
}

func CreateJWT(data Payload, secret string) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	//header converted to bytearray
	byte_arr_header, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	//bytearray to base64
	base64_header := base64encoding(byte_arr_header)

	//payload converted to bytearray
	byte_arr_payload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	//bytearray to base64
	base64_payload := base64encoding(byte_arr_payload)

	//concat payload and header = message
	message := base64_header + "." + base64_payload

	//message to bytearray
	byte_arr_message := []byte(message)

	//secret to bytesecret
	byte_arr_secret := []byte(secret)

	h := hmac.New(sha256.New, byte_arr_secret)
	h.Write(byte_arr_message)

	signature := h.Sum(nil)

	base64_signature := base64encoding(signature)

	jwt := base64_header + "." + base64_payload + "." + base64_signature

	return jwt, nil

}

func base64encoding(bytearray []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(bytearray)
}

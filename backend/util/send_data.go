package util

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, data interface{}, status int) {
	json.NewEncoder(w).Encode(data)
}

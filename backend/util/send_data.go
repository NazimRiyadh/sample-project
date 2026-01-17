package util

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, data interface{}, r *http.Request) {
	json.NewEncoder(w).Encode(data)
}

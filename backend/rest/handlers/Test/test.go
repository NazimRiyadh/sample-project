package Test

import (
	"fmt"
	"net/http"
)

func (h *TestHandler) Test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am serving inside test")
}

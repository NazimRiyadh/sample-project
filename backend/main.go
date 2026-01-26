package main

import (
	"ecommerce/util"
	"fmt"
)

func main() {
	//cmd.Serve()
	jwt, err := util.CreateJWT(util.Payload{
		ID:        1,
		Firstname: "riyadh",
		Lastname:  "nazim",
		Email:     "[EMAIL_ADDRESS]",
		IsOwner:   true,
	}, "secret")
	fmt.Println(jwt)
	fmt.Println(err)
}

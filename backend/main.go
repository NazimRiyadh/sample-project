package main

import (
	"fmt"
)

type iuser interface {
	PrintDetails()
	//RecieveMoney(amount float64) float64
}

type user struct {
	name  string
	age   int
	class int
	roll  int
}

func (obj user) PrintDetails() {
	fmt.Println("yoyo")
}

func main() {
	//cmd.Serve()
	var Youser iuser
	Youser = user{
		name:  "riyadh",
		age:   12,
		class: 6,
		roll:  10,
	}
	Youser.PrintDetails()
	fmt.Println(Youser)
}

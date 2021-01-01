package main

import "fmt"

type money float32

type dollar money

func main() {

	var rupee money = 20
	var sum dollar = 30
	var i float32 = float32(rupee)
	_ = i
	_ = sum
	fmt.Println(rupee)

}

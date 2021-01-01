package main

import "fmt"

func main() {

	defer fmt.Println("Hello")
	fmt.Println("abc")
	panic("I want to exit")
	// defer fmt.Println("Hyyyyy")
	fmt.Println("Bye")

}

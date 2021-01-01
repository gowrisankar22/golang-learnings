package main

import "fmt"

func main() {

	show("abc", 4, 87, 8, 9, 65, 56)
}

func show(name string, a ...int) {
	fmt.Printf("%T", a)
	fmt.Println(a)
}

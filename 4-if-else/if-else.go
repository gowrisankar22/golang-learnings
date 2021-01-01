package main

import "fmt"

func main() {
	a, b, c := 10, 300, 20

	if a > b && a > c {
		fmt.Println("A")
	} else if b > a && b > c {
		fmt.Println("B")
	} else {
		fmt.Println("c")
	}

	ok := true
	if ok {
		//
	}

}

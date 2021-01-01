package main

import "fmt"

func main() {
	a := []int{10, 30, 40, 50}

	//b := make([]int, 3)
	//
	//copy(b, a)

	b := make([]int, len(a), cap(a))
	copy(b, a)
	fmt.Println(a, b)

}

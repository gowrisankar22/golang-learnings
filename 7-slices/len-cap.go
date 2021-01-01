package main

import "fmt"

func main() {
	var i []int
	//i[0] = 1

	i = append(i, 10, 20, 30) // [10,20,30,0,0] //backing array
	fmt.Printf("1st Time %p\n", i)
	fmt.Println(len(i), cap(i))
}

package main

import "fmt"

func main() {

	var i []int
	i = make([]int, 2, 5)

	//x:=make([]int, 2, 5)

	i[0] = 10
	fmt.Printf("1st Time %p\n", i)

	i = append(i, 40)
	fmt.Printf("2nd Time %p\n", i)

	fmt.Println(i)

}

package main

import "fmt"

func main() {
	var i []int
	//i[0] = 1

	i = append(i, 10, 20, 30) // [10,20,30,0,0] //backing array
	fmt.Printf("1st Time %p\n", i)
	i = append(i, 40)
	fmt.Printf("2nd Time %p\n", i)

	i = append(i, 50)
	fmt.Printf("3rd Time %p\n", i)

	fmt.Printf("%#v", i)
}

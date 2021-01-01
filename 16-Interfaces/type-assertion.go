package main

import "fmt"

func main() {
	var i interface{} = 2
	a, ok := i.(int)

	fmt.Println(a,ok)

}

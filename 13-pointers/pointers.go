package main

import "fmt"

func main() {

	var i int = 10  //123 -> a -> b ->c
	var x *int = &i //123 -> b -> c

	fmt.Printf("%p\n", &i)
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", *x)

	*x = 20
	i = 30

	fmt.Println("i=", i)
	fmt.Printf("x=%v\n", *x)

}

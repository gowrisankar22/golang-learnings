package main

import "fmt"

func main() {
	var i, v interface{} = "Hello", true
	// interface is an abstract type

	fmt.Printf("%T %T\n", i, v)
	disp(i, v)
}

func disp(a ...interface{}) {
	fmt.Printf("%#v\n", a)
	fmt.Printf("%T\n", a[0])

}

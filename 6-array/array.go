package main

import "fmt"

func main() {
	//
	//a := [5]int{2, 3, 5, 7} // last value will be 0 as int default value is 0
	//fmt.Printf("%#v", a)

	var a [5]string

	a[0] = "Some Value"
	fmt.Printf("%#v", a)

	for i, v := range a {
		fmt.Println(i, v)
	}

}

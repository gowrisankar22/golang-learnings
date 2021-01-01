package main

import "fmt"

func main() {

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	z := 0
	for ; z <= 5; z++ {
		//
	}

	c := 0
	for c <= 5 {
		//
		if c == 2 {
			break
		}
		c++
	}

}

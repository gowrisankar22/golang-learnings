package main

import "fmt"

func main() {

	a := [2][3]int{
		{3, 65},
		{54, 7, 9},
	}

	for i, v := range a {
		fmt.Println(i, v)
		for i, val := range v {
			fmt.Println(i, val)
		}
	}

}

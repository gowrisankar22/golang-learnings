package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	p := &person{
		name: "abc",
		age:  20,
	}

	fmt.Println(p.name)

}

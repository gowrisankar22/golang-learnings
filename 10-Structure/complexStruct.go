package main

import "fmt"

type names []string

type class struct {
	name          names
	totalStudents int
	class         int
}

func main() {

	c1 := class{
		name:          names{"Rahul", "Ajay"},
		totalStudents: 2,
		class:         1,
	}
	fmt.Println(c1)
	for i, v := range c1.name {
		fmt.Println(i,v)
	}

}

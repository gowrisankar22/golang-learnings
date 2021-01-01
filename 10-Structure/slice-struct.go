package main

import "fmt"

type Student struct {
	name string
	age  int
}
type Students []Student

func main() {
	students := Students{ //[]Students
		{
			name: "abc",
			age:  20,
		},
		{
			name: "xyz",
			age:  24,
		},
	}

	for _, v := range students {
		fmt.Printf("%T  %v \n", v, v.name)
	}
	students[0].age = 33
	fmt.Println(students[0].name)

}

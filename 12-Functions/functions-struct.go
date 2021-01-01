package main

import "fmt"

type user struct {
	name string
	age  int
}

func UpdateValues(u user) user { // copied
	u.name = "dev"
	u.age = 30
	return u
}

func arrays(a []int) {
	fmt.Println(a)
}

func main() {

	u := user{
		name: "abc",
		age:  10,
	}

	u = UpdateValues(u)
	fmt.Println(u)

	//i:= []int {1,2,3}
	i:= [5]int {1,2,3}
	arrays(i[:])

}

package main

import "fmt"

type user struct {
	name string
	age  int
}

func UpdateStruct(u *user) {
	u.name = "dev"
	u.age = 18
}

func main() {

	u := user{
		name: "xyz",
		age:  24,
	}

	UpdateStruct(&u)
	fmt.Println(u)
}

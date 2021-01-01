package main

import "fmt"

type user struct {
	name string
	age  int
}

func (u user) show(name string) {
	u.name = name
	fmt.Println(u.name, u.age)
}
func (u user) update() error {

	return nil
}

func main() {
	u1 := user{
		name: "abc",
		age:  20,
	}
	u1.show("ajay")
	err := u1.update()
	_=err
}

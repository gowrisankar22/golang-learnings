package main

import "fmt"

type student struct {
	name string
}

type class struct {
	ClassName string
}

func (s student) show() {
	fmt.Println(s.name)
}

func (c class) show() {
	fmt.Println(c)
}

func main() {
	s := student{name: "abc"}
	c := class{ClassName: "1"}
	s.show()
	c.show()

}

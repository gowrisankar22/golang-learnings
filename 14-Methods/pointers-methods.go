package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) fillValues(name string, age int) {
	p.name = name
	p.age = age
}

func main() {
	p := person{}
	p.fillValues("xyz", 20)

	fmt.Println(p)
}

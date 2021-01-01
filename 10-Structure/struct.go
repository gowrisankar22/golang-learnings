package main

import "fmt"

type user struct {
	name string //fields
	age  int
	add  string
}

func main() {
	//var student struct {
	//	name string
	//}
	//var student1 struct {
	//	name string
	//}
	//var studen2 struct {
	//	name string
	//}

	//student.name = "hello"

	var u user
	u.name = "Dev"
	u.age = 10
	u.add = "abc"

	fmt.Printf("%#v\n", u)
	fmt.Printf("%#v\n", u.age)

	rahul := user{
		name: "Rahul",
		age:  20,
		add:  "abc",
	}

	rahul.age = 25
	fmt.Println(rahul)

}

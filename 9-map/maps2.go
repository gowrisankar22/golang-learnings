package main

import "fmt"

func main() {

	var student = make(map[int][]string)
	student[1] = []string{"abc", "xyz"}
	student[2] = []string{"zxc", "gyz"}

	fmt.Printf("%#v", student)

}

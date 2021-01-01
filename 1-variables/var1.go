package main

import "fmt"

func main() {
	var (
		name  string
		age   int     = 5
		marks float32 = 6
	)
	name = "Dev"

	var abc = "abc"

	var (
		a, b, c, d int
		e          string
	)

	_, _, _, _, _ = a, b, c, d, e
	_ = abc

	fmt.Println(name, age, marks)
}

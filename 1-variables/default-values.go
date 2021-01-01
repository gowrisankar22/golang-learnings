package main

import "fmt"

func main() {
	var (
		a  int
		b  float32
		s  string
		ok bool
	)

	fmt.Println(a, b, s, ok)
	fmt.Printf("int= %v str=%q \n", a, s)
}

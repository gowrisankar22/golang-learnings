package main

import "fmt"

func main() {
	/*
		a=b

		b=append(b,"Data")


	*/

	a := []int64{12, 45, 67, 98, 78, 89, 65, 90, 55}

	b := a[1:6]
	fmt.Println("cap a ", cap(a), "len a", len(a))

	fmt.Println("cap b ", cap(b), "len b", len(b))

	b = append(b, 109)

	fmt.Printf("a %p b %p", a, b)

	fmt.Println("\na=", a, "b=", b)
}

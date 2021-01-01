package main

import "fmt"

func updateValues(i *int) {
	*i = 20
}

func main() {
	var a int
	updateValues(&a)
	fmt.Println(a)

}

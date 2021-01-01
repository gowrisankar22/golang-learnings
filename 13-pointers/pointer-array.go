package main

import "fmt"

func fillValues(a *[5]int) {
	j := 0
	for i := 5; i > 0; i-- {

		a[j] = i
		j++
	}
}

func main() {
	var a [5]int
	fillValues(&a)
	fmt.Println(a)


}

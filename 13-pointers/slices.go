package main

import "fmt"

func slices(a []int) {
	j := 0
	for i := 5; i > 0; i-- {
		a[j] = i
		j++
	}

}

func main() {

	b := make([]int, 5, 5)
	slices(b)

	fmt.Println("From slices", b)
}

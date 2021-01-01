package main

import (
	"fmt"
	"strconv"
)

func add(a int, s string) (int, error) {
	b, err := strconv.Atoi(s)

	if err != nil {
		return 0, err
	}
	return a + b, nil

}

func main() {

	fmt.Println(add(2, "abc"))

}

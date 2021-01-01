package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
// i+=2  // i = i+2  // i*=2 // i = i*2
}

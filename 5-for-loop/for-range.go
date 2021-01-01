package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args

	for i, v := range args {
		fmt.Println(i, v)

		fmt.Println(args[i])
		if i == 0 {
			break
		}
	}

	//for _, v := range args {
	//	fmt.Println(v)
	//}

	//for i := range args {
	//	fmt.Println(i)
	//}

}

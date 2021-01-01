package main

import "fmt"

func main() {

	CheckType(2)

}

func CheckType(i interface{}) {

	switch v := i.(type) {

	case int:
		fmt.Println("Int", v)
	case string:
		fmt.Println("String", v)
	case bool:
		fmt.Println("bool", v)
	default:
		fmt.Println("Cannot identify", v)
	}

}

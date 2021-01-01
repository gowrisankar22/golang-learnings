package main

import "fmt"

func main() {

	var day string

	fmt.Println("Enter day name")
	fmt.Scanln(&day)

	switch day {
	case "monday":
		fmt.Println("Monday")
		fallthrough

	case "tuesday":
		fmt.Println("Tuesday")
		fallthrough

	default:
		fmt.Println("Sorry")
	}

}

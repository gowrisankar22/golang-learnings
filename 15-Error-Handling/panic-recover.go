package main

import (
	"fmt"
	"os"
)

func open(filename string) (*os.File, error) {

	_, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	return nil, nil

}

func recoverFileNotFound() {
	r := recover()

	if r != nil {
		fmt.Println("recovered", r)
	}

}

func abc() {
	defer recoverFileNotFound()
	_, err := open("abc")
	if err != nil {
		panic(err)
	}
}

func main() {


	abc()

	fmt.Println("fine")
}

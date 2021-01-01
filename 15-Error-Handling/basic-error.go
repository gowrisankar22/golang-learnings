package main

import (
	"fmt"
	"log"
	"os"
)

func openFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)

	if err != nil {

		return nil, err
	}
	return file, nil

}

func main() {

	_, err := openFile("abc.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("fine")

}

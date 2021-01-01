package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var fileNotFound error = errors.New("file not found abc.txt")

func openFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)

	if err != nil {
		//customErr := fmt.Errorf("%v %v", err, fileNotFound) // go 1.13
		return nil, fmt.Errorf("%v %v", err, fileNotFound)
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

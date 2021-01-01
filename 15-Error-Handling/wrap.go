package main

import (
	"errors"
	"fmt"
	"os"
)

var fileNotFound error = errors.New("file not found abc.txt")

func openFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, fmt.Errorf("%v %w ", err, fileNotFound)
	}
	return file, nil

}

func main() {

	_, err := openFile("abc.txt")

	if errors.Is(err, fileNotFound) { // if err == customErr {}
		fmt.Println("match")
	} else {
		fmt.Println("not")
	}
}

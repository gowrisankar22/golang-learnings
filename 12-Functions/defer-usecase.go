package main

import "os"

func main() {

	file, err := os.OpenFile("test", os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		if os.IsPermission(err) {
			panic(err)
		}

	}

}

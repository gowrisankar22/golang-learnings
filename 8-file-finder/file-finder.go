package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("Provide a directory name")
	}

	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		log.Fatalf("%v", err)
	}

	var names []byte
	for _, f := range files {
		//fmt.Printf("%#v\n\n", f)

		if f.Size() == 0 {
			names = append(names, f.Name()...)
			names = append(names, '\n')
		}
	}

	err = ioutil.WriteFile("out.txt", names, 666)

	if err != nil {
		log.Fatal(err)
	}

}

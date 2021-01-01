package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {

	for i := 1; i <= 5; i++ {
		wg.Add(1)   //5
		go hello(i) // 5 // 5 goroutines
	}
	fmt.Println("Main")

	wg.Wait()
}

func hello(i int) {
	fmt.Println("Hello", i)
	wg.Done()
}

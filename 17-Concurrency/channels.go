package main

import (
	"fmt"
	"sync"
)

var w = sync.WaitGroup{}

func main() {

	c := make(chan int)

	w.Add(4)
	go add(2, 2, c)
	go sub(4, 2, c)
	go multiply(3, 3, c)
	go calc(c)

	w.Wait()

}

func add(a int, b int, c chan int) {
	fmt.Println("exec add")
	sum := a + b
	c <- sum

	w.Done()
}

func sub(a int, b int, c chan int) {
	fmt.Println("exec sub")
	sum := a - b
	c <- sum
	w.Done()

}

func multiply(a int, b int, c chan int) {
	fmt.Println("exec multiply")
	sum := a * b
	c <- sum
	w.Done()
}

func calc(c chan int) {
	fmt.Println("exec calc")
	a, b, d := <-c, <-c, <-c
	fmt.Println(a + b + d)
	w.Done()

}

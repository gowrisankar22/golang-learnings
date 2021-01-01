package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan string)
	c1 := make(chan string)
	c3 := make(chan string)

	go func() {
		time.Sleep(10 * time.Second)
		c <- "one"

	}()

	go func() {
		time.Sleep(10 * time.Second)
		c1 <- "two"

	}()

	go func() {
		c3 <- "three"
	}()

	a := <-c
	fmt.Println("received ", a)

	b := <-c1
	fmt.Println("received ", b)

	d := <-c3
	fmt.Println("received ", d)

}

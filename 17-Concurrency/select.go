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
		time.Sleep(5 * time.Second)
		c <- "one"
	}()

	go func() {
		c1 <- "two"
	}()

	go func() {
		c3 <- "three"
	}()

	for i := 1; i <= 3; i++ {
		select {
		case a := <-c:
			fmt.Println("received ", a)

		case b := <-c1:
			fmt.Println("received ", b)

		case c := <-c3:
			fmt.Println("received ", c)
		}
	}
}

package main

import (
	"fmt"
	"time"
)

func main() {

	go disp()
	time.Sleep(5 * time.Millisecond)
	fmt.Println("Hello from main")
	time.Sleep(5 * time.Millisecond)

}

func disp() {
	//panic("panic")

	fmt.Println("Hello from disp")

}

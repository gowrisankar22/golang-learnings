package main

import (
	"fmt"
	"sync"
	"time"
)

var cabs = 2

var wg = sync.WaitGroup{}

func main() {
	m := &sync.Mutex{}
	names := []string{"Ravi", "Raj", "Dev", "Ankit", "Vipin"}

	for _, n := range names {
		wg.Add(1)
		go cab(n, m)
	}
	wg.Wait()
}

func cab(name string, m *sync.Mutex) {

	fmt.Println("Welcome to our site")

	m.Lock()
	if cabs >= 1 {

		fmt.Println("Cab is available for ", name)
		time.Sleep(1 * time.Second)
		fmt.Println("Booking is confirmed for", name)
		fmt.Println("Thanks ", name)
		cabs--
	} else {
		fmt.Println("Cab is not available for ", name)
	}
	m.Unlock()
	wg.Done()

}

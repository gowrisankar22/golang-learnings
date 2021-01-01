package main

import "fmt"

type shape interface {
	Area() float64
	Perimeter() float64
}

type rect struct {
	width  float64
	height float64
}

func (r rect) Area() float64 {
	return r.height * r.width
}

func (r rect) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (r rect) hello() {
	fmt.Println("i am hello")
}

type square struct {
	width  float64
	height float64
}

func (s square) Area() float64 {
	return 2 * 2
}
func (s square) Perimeter() float64 {
	return 4 * 4
}

func main() {

	var s shape
	r := rect{
		width:  10,
		height: 10,
	}
	sq := square{}

	s = r

	fmt.Printf("%v\n for rect", s == r)
	fmt.Printf("%T\n", s)

	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())

	s = sq

	fmt.Printf("%v\n for square ", s == r)
	fmt.Printf("%T\n", s)

	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())

}

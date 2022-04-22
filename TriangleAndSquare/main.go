package main

import "fmt"

func main() {
	t := triangle{base: 10, height: 30}
	s := square{sideLength: 10}

	fmt.Println("Triangle ", t.getArea())
	fmt.Println("Square ", s.getArea())

}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println("Area: ", s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

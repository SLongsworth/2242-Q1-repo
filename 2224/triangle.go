package main

import "fmt"

type Triangle struct {
	base       float64
	height     float64
	hypotenuse float64
}

func (t Triangle) area() float64 {
	return t.base * t.height / 2
}
func (t Triangle) perimeter() float64 {
	return t.base + t.height + t.hypotenuse
}

func main() {
	t := Triangle{
		base:       4,
		height:     3,
		hypotenuse: 5}

	fmt.Println("Area: ", t.area())
	fmt.Println("Perimeter:", t.perimeter())
}

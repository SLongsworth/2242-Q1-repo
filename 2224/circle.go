package main

import "fmt"

type Circle struct {
	radius float64
}

func (c *Circle) area() float64 {
	return (22 / 7) * (c.radius * c.radius)
}

func (c Circle) perimeter() float64 {
	return (2) * (22 / 7) * (c.radius)
}

func main() {
	c := Circle{
		radius: 10,
	}
	fmt.Println("Area: ", c.area())
	fmt.Println("Perimeter: ", c.perimeter)
}

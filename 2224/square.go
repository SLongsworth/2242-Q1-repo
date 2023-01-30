package main

import "fmt"

type Square struct {
	side float64
}

func (s *Square) area() float64 {
	return s.side * s.side
}

func (s Square) perimeter() float64 {
	return (4) * s.side
}

func main() {
	s := Square{
		side: 9,
	}
	fmt.Println("Area: ", s.area())
	fmt.Println("Perimeter: ", s.perimeter())
}

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() error
}

type Circle struct {
	radius float64
}

type Square struct {
	side int
}

func (s *Square) Area() error {
	fmt.Println("The area of the square is : ", s.side*s.side)
	return nil
}
func (s *Circle) Area() error {
	fmt.Println("The radius of the circle is :", math.Pi*math.Pow(s.radius, 2))
	return nil
}

func ProcessArea(s Shape) error {
	return s.Area()
}

func main() {
	err := ProcessArea(&Circle{
		radius: 3,
	})
	if err != nil {
		return
	}

	err = ProcessArea(&Square{
		side: 3,
	})
	if err != nil {
		return
	}
}

// Rectangle bad code
type Rectangle struct {
	height int
	width  int
}

func Area(r Rectangle) {
	fmt.Println("Area of rectangle is :", r.height*r.width)
}

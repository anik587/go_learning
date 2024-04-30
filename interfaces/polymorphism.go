package interfaces

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	circumf() float64
}

type Square struct {
	Length float64
}
type Circle struct {
	Radius float64
}

// square methods
func (s Square) area() float64 {
	return s.Length * s.Length
}
func (s Square) circumf() float64 {
	return s.Length * 4
}

func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) circumf() float64 {
	return 2 * math.Pi * c.Radius
}

func PrintShapeInfo(s Shape) {
	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("circumference of %T is: %0.2f \n", s, s.circumf())
}

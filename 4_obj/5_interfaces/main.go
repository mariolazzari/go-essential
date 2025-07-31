package main

import (
	"fmt"
	"math"
)

type Square struct {
	Length float64
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type Shape interface {
	Area() float64
}

func sumAreas(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func main() {
	s1 := Square{Length: 5}
	c1 := Circle{Radius: 3}
	s2 := Square{Length: 6}
	c2 := Circle{Radius: 4}

	shapes := []Shape{s1, c1, s2, c2}

	totalArea := sumAreas(shapes)
	fmt.Printf("Total Area: %f\n", totalArea)
}

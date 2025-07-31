package main

import "fmt"

type Square struct {
	X      int
	Y      int
	Length int
}

func NewSquare(x, y, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("Side length must be positive")
	}

	s := Square{X: x, Y: y, Length: length}

	return &s, nil
}

func (s Square) Area() int {
	return s.Length * s.Length
}

func (s Square) Perimeter() int {
	return 4 * s.Length
}

func (s *Square) Move(x, y int) {
	s.X += x
	s.Y += y
}

func main() {

	s, err := NewSquare(0, 0, 5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	println("Area:", s.Area())
	println("Perimeter:", s.Perimeter())
	s.Move(2, 3)
	println("Moved to:", s.X, s.Y)

	s, err = NewSquare(0, 0, -5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	println("Area:", s.Area())
	println("Perimeter:", s.Perimeter())
}

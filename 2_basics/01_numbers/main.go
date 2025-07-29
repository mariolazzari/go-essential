package main

import "fmt"

func main() {
	/*
		var x int
		var y int

		x = 1
		y = 2
	*/

	// Declare and initialize variables
	x, y := 1.0, 2.0

	fmt.Printf("x = %v, of type %T\n", x, x)
	fmt.Printf("y = %v, of type %T\n", y, y)

	// type inference :=
	mean := (x + y) / 2.0
	fmt.Printf("The mean of %v and %v is: %v\n", x, y, mean)
}

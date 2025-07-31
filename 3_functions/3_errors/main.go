package main

import (
	"fmt"
	"math"
)

func main() {
	result, err := sqrt(16)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Square root of 16 is", result)
	}

	result, err = sqrt(-4)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Square root of -4 is", result)
	}
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0.0, fmt.Errorf("cannot take square root of negative number: %f", n)
	}

	return math.Sqrt(n), nil
}

package sqrt

import (
	"errors"
)

// Common errors
var (
	ErrNegSqrt    = errors.New("sqrt of negative number")
	ErrNoSolution = errors.New("no solution found")
)

// Abs returns the absolute value of val
func Abs(val float64) float64 {
	if val < 0 {
		return -val
	}
	return val
}

// Sqrt return the square root of a number
func Sqrt(val float64) (float64, error) {
	if val < 0.0 {
		return 0.0, ErrNegSqrt
	}

	if val == 0.0 {
		return 0.0, nil // shortcut
	}

	guess, epsilon := 1.0, 0.00001
	for range 10000 {
		if Abs(guess*guess-val) <= epsilon {
			return guess, nil
		}
		guess = (val/guess + guess) / 2.0
	}

	return 0.0, ErrNoSolution
}

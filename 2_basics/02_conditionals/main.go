package main

import "fmt"

func main() {
	x := 10

	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is not greater than 5")
	}

	if x > 0 && x <= 10 {
		fmt.Println("x is between 0 and 10")
	} else {
		fmt.Println("x is not between 0 and 10")
	}

	if x < 20 || x > 30 {
		fmt.Println("x is either less than 20 or greater than 30")
	} else {
		fmt.Println("x is between 20 and 30")
	}

	switch x {
	case 10:
		fmt.Println("x is exactly 10")
	case 5:
		fmt.Println("x is exactly 5")
	default:
		fmt.Println("x is neither 10 nor 5")
	}

	switch {
	case x > 5:
		fmt.Println("x is greater than 5")
	case x < 5:
		fmt.Println("x is less than 5")
	default:
		fmt.Println("x is equal to 5")
	}
}

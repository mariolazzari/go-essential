package main

import "fmt"

func main() {
	for i := range 3 {
		fmt.Println(i)
	}

	// Using a for loop with a condition
	fmt.Println("----")
	for i := range 3 {
		if i > 1 {
			break
		}
		fmt.Println(i)
	}

	// Using a for loop with continue
	fmt.Println("----")
	for i := range 3 {
		if i > 1 {
			continue
		}
		fmt.Println(i)
	}

	// Using a for loop with a condition and no initialization or post statement
	fmt.Println("----")
	a := 0
	for a < 3 {
		fmt.Println(a)
		a++
	}

	// Using a for loop with an infinite loop and break condition
	fmt.Println("----")
	for {
		fmt.Println("Infinite loop", a)
		a++
		if a > 5 {
			break
		}
	}

}

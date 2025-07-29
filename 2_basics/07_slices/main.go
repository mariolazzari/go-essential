package main

import "fmt"

func main() {
	// Declare and initialize an array of strings
	loons := []string{"loon1", "loon2", "loon3"}
	fmt.Printf("loons: %v, (type %T)\n", loons, loons)
	fmt.Println("Array length:", len(loons))
	fmt.Println("Array capacity:", cap(loons))
	fmt.Println("First element:", loons[0])

	// slices
	fmt.Println("Slicing the array:", loons[1:])

	for i, v := range loons {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}
	for _, v := range loons {
		fmt.Printf("Value: %s\n", v)
	}

	// Append to the slice
	loons = append(loons, "loon4")
	fmt.Printf("After appending: %v\n", loons)
}

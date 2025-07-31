package main

import "fmt"

func main() {
	// Demonstrating pass by reference and pass by value in Go
	nums := []int{1, 2, 3, 4, 5}
	doubleAt(nums, 2)
	fmt.Println(nums)

	// Demonstrating pass by value
	num := 10
	double(num)
	fmt.Println(num)

	// Demonstrating pass by reference with a pointer
	doublePtr(&num)
	fmt.Println(num)
}

// pass by reference
func doubleAt(nums []int, index int) {
	if index < 0 || index >= len(nums) {
		fmt.Println("Index out of bounds")
		return
	}

	nums[index] *= 2
}

// pass by value
func double(num int) {
	num *= 2
}

// pass by reference with a pointer
func doublePtr(num *int) {
	*num *= 2
}

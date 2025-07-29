package main

func main() {
	// Declare and initialize an array of integers
	numbers := []int{10, 200, 30, 400, 50}
	maxValue := max(numbers)
	println("The maximum value in the slice is:", maxValue)
}

func max(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	maxVal := nums[0]
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}

	return maxVal
}

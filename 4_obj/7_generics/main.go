package main

import "fmt"

type Ordered interface {
	int | float64 | string
}

func min[T Ordered](values []T) (T, error) {
	if len(values) == 0 {
		var zero T
		return zero, fmt.Errorf("cannot find minimum of an empty slice")
	}

	minValue := values[0]
	for _, value := range values[1:] {
		if value < minValue {
			minValue = value
		}
	}

	return minValue, nil
}

func main() {
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	minNum, err := min(nums)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Minimum number:", minNum)
	}

	strs := []string{"apple", "banana", "cherry"}
	minStr, err := min(strs)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Minimum string:", minStr)
	}
}

package main

import "fmt"

func main() {
	vals := []int{1, 2, 3}
	v, err := safeValue(vals, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Value:", v)
	}
	fmt.Println("Program continues after handling the error.")
}

func safeValue(vals []int, index int) (n int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("index out of range: %v", e)
		}
	}()

	return vals[index], err
}

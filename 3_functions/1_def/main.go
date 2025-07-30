package main

func main() {
	res := add(10, 20)
	println("Sum:", res)
	res, rem := divide(10, 3)
	println("Quotient:", res, "Remainder:", rem)
	res, rem = divide(10, 0)
	println("Quotient:", res, "Remainder:", rem)
}

func add(a int, b int) int {
	return a + b
}

func divide(a int, b int) (int, int) {
	if b == 0 {
		return 0, 0
	}
	return a / b, a % b
}

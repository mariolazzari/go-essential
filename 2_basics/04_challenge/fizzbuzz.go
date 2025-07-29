// print number between 1 and 20
// if number is divisible by 3 print "Fizz"
// if number is divisible by 5 print "Buzz"
// if number is divisible by 3 and 5 print "FizzBuzz"
// if number is not divisible by 3 or 5 print the number itself

package main

func main() {
	for i := 1; i <= 20; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			println("FizzBuzz")
		case i%3 == 0:
			println("Fizz")
		case i%5 == 0:
			println("Buzz")
		default:
			println(i)
		}
	}
}

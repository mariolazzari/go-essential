# Go Essentials: Concurrency, Connectivity, and High-Performance Apps

## Getting started

### Go program

```go
package main

import "fmt"

func main() {
    fmt.Println("Welcome to the Go programming world!")
}
```

### Go tool

```sh
go run welcome.go
go build welcome.go
go help
```

## Go basics

### Numbers and assignemnts

```go
package main

import "fmt"

func main() {
 /*
  var x int
  var y int

  x = 1
  y = 2
 */

 // Declare and initialize variables
 x, y := 1.0, 2.0

 fmt.Printf("x = %v, of type %T\n", x, x)
 fmt.Printf("y = %v, of type %T\n", y, y)

 // type inference :=
 mean := (x + y) / 2.0
 fmt.Printf("The mean of %v and %v is: %v\n", x, y, mean)
}
```

### Conditionals

```go
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
```

### For loop

```go
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
```

### Challenge: FizzBuzz

```go
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
```

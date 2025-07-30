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

### Strings

```go
package main

import "fmt"

func main() {
 // strings are immutable sequences of bytes
 book := "Delitto e Castigo"
 fmt.Println(book)

 // Print the first character
 fmt.Println("Length:", len(book))

 // Note: This will print the byte length, not the rune count
 fmt.Printf("First character: %v (type %T)", book[0], book[0])

 // slices are views into the string
 fmt.Println("\nSlice of the first 5 characters:", book[4:11])
 fmt.Println("Slice of the first 5 characters:", book[:5])
 fmt.Println("Slice of the last 5 characters:", book[len(book)-5:])

 // string concatenation
 author := "Fëdor Dostoevskij"
 fmt.Println("Concatenated string:", book+" by "+author)

 // multiline strings
 multiline := `This is a multiline string.
It can span multiple lines and includes "quotes" without escaping.
It is useful for longer text blocks.`
 fmt.Println("\nMultiline string:")
 fmt.Println(multiline)
}
```

### Even numbers

```go
// even ended number:
// A number is even ended if it ends with an even digit (0, 2, 4, 6, or 8).

package main

import "fmt"

const START = 1000
const END = 10000

func main() {
 count := 0
 for i := START; i < END; i++ {
  for j := START; j < END; j++ {
   n := i * j
   s := fmt.Sprintf("%d", n)
   if s[0] == s[len(s)-1] {
    count++
   }
  }
 }

 fmt.Printf("Count of even ended numbers form %d to %d: %d\n", START, END-1, count)
}
```

### Slices

```go
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
```

### Challenge: max value

```go
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
```

### Maps

```go
package main

import "fmt"

func main() {
 stocks := map[string]float64{
  "GOOGL": 2800.00,
  "AAPL":  150.00,
  "AMZN":  3400.00,
  "MSFT":  299.00,
 }

 // map length
 fmt.Println(len(stocks))

 // print value
 fmt.Println(stocks["AAPL"])
 // zero value for non-existing key
 fmt.Println(stocks["TSLA"])
 // check if key exists
 price, ok := stocks["TSLA"]
 if ok {
  fmt.Println("TSLA exists with price:", price)
 } else {
  fmt.Println("TSLA does not exist in the map.")
 }

 // set value
 stocks["TSLA"] = 800.00
 fmt.Println("After adding TSLA:", stocks)

 // delete value
 delete(stocks, "AMZN")
 fmt.Println("After deleting AMZN:", stocks)

 // print all keys
 for key := range stocks {
  fmt.Println(key)
 }

 // print all keys and values
 for key, value := range stocks {
  fmt.Println(key, ":", value)
 }
}
```

### Challenge: words counter

```go
package main

import "strings"

func main() {
 text := "Hello world! This is a test. Hello world again as a second test."
 wordsCount := map[string]int32{}

 works := strings.SplitSeq(text, " ")
 for word := range works {
  word = strings.ToLower(word)
  wordsCount[word]++
 }

 for word, count := range wordsCount {
  println(word, ":", count)
 }
}
```

## Functions

### Function definition

```go
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
```

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

### Parameter passing

```go
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
```

### Error return

```go
package main

import (
 "fmt"
 "math"
)

func main() {
 result, err := sqrt(16)
 if err != nil {
  fmt.Println("Error:", err)
 } else {
  fmt.Println("Square root of 16 is", result)
 }

 result, err = sqrt(-4)
 if err != nil {
  fmt.Println("Error:", err)
 } else {
  fmt.Println("Square root of -4 is", result)
 }
}

func sqrt(n float64) (float64, error) {
 if n < 0 {
  return 0.0, fmt.Errorf("cannot take square root of negative number: %f", n)
 }

 return math.Sqrt(n), nil
}
```

### Defer

```go
package main

import "fmt"

func main() {
 worker()
}

func worker() {
 r1, err := acquire("resource1")
 if err != nil {
  fmt.Println("Error acquiring resource1:", err)
  return
 }
 defer release(r1)

 r2, err := acquire("resource2")
 if err != nil {
  fmt.Println("Error acquiring resource2:", err)
  return
 }
 defer release(r2)

 fmt.Println("Working with", r1)
 fmt.Println("Working with", r2)
 fmt.Println("Done working")
}

func acquire(name string) (string, error) {
 return name, nil
}

func release(name string) {
 fmt.Println("Releasing", name)
}
```

### Challenge

```go
package main

import (
 "fmt"
 "net/http"
)

func main() {
 ctype, err := contentType("https://mariolazzari.it")
 if err != nil {
  fmt.Println("Error:", err)
 } else {
  fmt.Println("Content-Type:", ctype)
 }
}

func contentType(url string) (string, error) {
 res, err := http.Get(url)
 if err != nil {
  return "", err
 }
 defer res.Body.Close()

 ctype := res.Header.Get("Content-Type")
 if ctype == "" {
  return "", fmt.Errorf("Content-Type header not found")
 }

 return ctype, nil
}
```

## Object oriented

### Structs

```go
package main

import (
 "fmt"
 "time"
)

type Budget struct {
 CampaignID string
 Balance    float64
 Expires    time.Time
}

func main() {
 // Creating a budget with all fields filled
 // The Expires field is set to a specific time in the future
 // Here we set it to 7 days from now
 b1 := Budget{
  CampaignID: "12345",
  Balance:    1000.00,
  Expires:    time.Now().Add(7 * 24 * time.Hour),
 }
 fmt.Println("Budget 1:", b1)
 fmt.Printf("%v\n", b1)

 // Creating a second budget without the Expires field
 // This will use the zero value for time.Time, which is the zero time
 // (January 1, year 1, 00:00:00 UTC)
 b2 := Budget{
  CampaignID: "67890",
  Balance:    500.00,
 }
 fmt.Println("Budget 2:", b2)
 fmt.Printf("%v\n", b2)

 // Creating a third budget without any fields set
 // This will use the zero values for all fields
 b3 := Budget{}
 fmt.Println("Budget 3:", b3)
 fmt.Printf("%v\n", b3)
}
```

### Methods

```go
package main

import (
 "fmt"
 "time"
)

type Budget struct {
 CampaignID string
 Balance    float64
 Expires    time.Time
}

func (b Budget) IsExpired() bool {
 return time.Now().After(b.Expires)
}

func (b Budget) TimeLeft() time.Duration {
 if b.IsExpired() {
  return 0
 }
 return b.Expires.Sub(time.Now().UTC())
}

// Update method to add a sum to the Balance
// This method modifies the Budget's Balance field
// It does not return a new Budget, but updates the existing one
// This allows for modifying the budget without creating a new instance
func (b *Budget) Update(sum float64) {
 b.Balance += sum
}

func main() {
 b := Budget{
  CampaignID: "12345",
  Balance:    1000.00,
  Expires:    time.Now().Add(7 * 24 * time.Hour),
 }
 fmt.Println("Time left:", b.TimeLeft())

 b.Update(500.00)
 fmt.Println("Updated Balance:", b.Balance)
}
```

### New function

```go
package main

import (
 "fmt"
 "time"
)

type Budget struct {
 CampaignID string
 Balance    float64
 Expires    time.Time
}

func NewBudget(campaignID string, balance float64, expires time.Time) (*Budget, error) {

 if campaignID == "" {
  return nil, fmt.Errorf("campaignID cannot be empty")
 }

 if balance < 0 {
  return nil, fmt.Errorf("balance cannot be negative")
 }

 if expires.Before(time.Now()) {
  return nil, fmt.Errorf("expires time must be in the future")
 }

 b := Budget{
  CampaignID: campaignID,
  Balance:    balance,
  Expires:    expires,
 }

 return &b, nil
}

func main() {
 expires := time.Now().Add(7 * 24 * time.Hour)
 b1, err := NewBudget("12345", 1000.00, expires)
 if err != nil {
  fmt.Println("Error creating budget:", err)
  return
 }
 fmt.Println("Budget created with CampaignID:", b1.CampaignID)
 fmt.Println("Initial Balance:", b1.Balance)
 fmt.Println("Expires at:", b1.Expires)

 b2, err := NewBudget("", 500.00, expires)
 if err != nil {
  fmt.Println("Error creating budget:", err)
 } else {
  fmt.Printf("Budget created with CampaignID: %s\n", b2.CampaignID)
 }
}
```

### Challenge structs

```go
package main

import "fmt"

type Square struct {
 X      int
 Y      int
 Length int
}

func NewSquare(x, y, length int) (*Square, error) {
 if length <= 0 {
  return nil, fmt.Errorf("Side length must be positive")
 }

 s := Square{X: x, Y: y, Length: length}

 return &s, nil
}

func (s Square) Area() int {
 return s.Length * s.Length
}

func (s Square) Perimeter() int {
 return 4 * s.Length
}

func (s *Square) Move(x, y int) {
 s.X += x
 s.Y += y
}

func main() {

 s, err := NewSquare(0, 0, 5)
 if err != nil {
  fmt.Println("Error:", err)
  return
 }

 println("Area:", s.Area())
 println("Perimeter:", s.Perimeter())
 s.Move(2, 3)
 println("Moved to:", s.X, s.Y)

 s, err = NewSquare(0, 0, -5)
 if err != nil {
  fmt.Println("Error:", err)
  return
 }
 println("Area:", s.Area())
 println("Perimeter:", s.Perimeter())
}
```

### Interfaces

```go
package main

import (
 "fmt"
 "math"
)

type Square struct {
 Length float64
}

func (s Square) Area() float64 {
 return s.Length * s.Length
}

type Circle struct {
 Radius float64
}

func (c Circle) Area() float64 {
 return math.Pi * math.Pow(c.Radius, 2)
}

type Shape interface {
 Area() float64
}

func sumAreas(shapes []Shape) float64 {
 total := 0.0
 for _, shape := range shapes {
  total += shape.Area()
 }
 return total
}

func main() {
 s1 := Square{Length: 5}
 c1 := Circle{Radius: 3}
 s2 := Square{Length: 6}
 c2 := Circle{Radius: 4}

 shapes := []Shape{s1, c1, s2, c2}

 totalArea := sumAreas(shapes)
 fmt.Printf("Total Area: %f\n", totalArea)
}
```

### Interfeces challenge

```go
package main

import (
 "fmt"
 "io"
 "os"
)

type Capper struct {
 w io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {
 for i, b := range p {
  if b >= 'a' && b <= 'z' {
   p[i] -= 32 // Convert to uppercase
  }
 }
 return c.w.Write(p)
}

func main() {
 c := &Capper{os.Stdout}
 fmt.Fprintf(c, "Hello, World")
}
```

### Generics

```go
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
```

## Errors

### pkg/errors

```go
// pkg/errors example
package main

import (
 "fmt"
 "log"
 "os"

 "github.com/pkg/errors"
)

// Config holds configuration
type Config struct {
 // configuration fields go here (redacted)
}

func readConfig(path string) (*Config, error) {
 file, err := os.Open(path)
 if err != nil {
  return nil, errors.Wrap(err, "can't open configuration file")
 }
 defer file.Close()

 cfg := &Config{}
 // Parse file here (redacted)
 return cfg, nil

}

func setupLogging() {
 out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
 if err != nil {
  return
 }
 log.SetOutput(out)
}

func main() {
 setupLogging()
 cfg, err := readConfig("/path/to/config.toml")
 if err != nil {
  fmt.Fprintf(os.Stderr, "error: %s\n", err)
  log.Printf("error: %+v", err)
  os.Exit(1)
 }

 // Normal operation (redacted)
 fmt.Println(cfg)
}
```

### Panic and recover

```go
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
```

### Challenge: process kill

```go
package main

import (
 "fmt"
 "os"
)

func main() {
 err := killServer("server.pid")
 if err != nil {
  fmt.Println("Error killing server:", err)
 } else {
  fmt.Println("Server killed successfully")
 }
 fmt.Println("Exiting main function")
}

func killServer(pidFile string) error {
 file, err := os.Open(pidFile)
 if err != nil {
  return fmt.Errorf("failed to read PID file: %w", err)
 }
 defer file.Close()

 var pid int
 if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
  return fmt.Errorf("failed to parse PID from file: %w", err)
 }

 fmt.Println("Killing server with PID:", pid)

 if err := os.Remove(pidFile); err != nil {
  return fmt.Errorf("failed to remove PID file: %w", err)
 }

 return nil
}
```

## Concurrency

### Goroutines

```go
package main

import (
 "fmt"
 "net/http"
 "sync"
 "time"
)

func main() {
 urls := []string{
  "https://www.google.com",
  "https://www.facebook.com",
  "https://www.twitter.com",
  "https://www.mariolazzari.it"}

 start := time.Now()
 siteSerial(urls)
 fmt.Println("Done with serial processing in ", time.Since(start))

 start = time.Now()
 siteConcurrent(urls)
 fmt.Println("Done with concurrent processing in ", time.Since(start))
 fmt.Println("Done with main")
}

func returnType(url string) {
 res, err := http.Get(url)
 if err != nil {
  fmt.Printf("Error fetching URL %s: %v\n", url, err)
  return
 }
 defer res.Body.Close()

 ctype := res.Header.Get("Content-Type")
 if ctype == "" {
  fmt.Printf("No Content-Type header found for URL %s\n", url)
  return
 }
 fmt.Printf("Content-Type for URL %s: %s\n", url, ctype)
}

func siteSerial(urls []string) {
 for _, url := range urls {
  returnType(url)
 }
}

func siteConcurrent(urls []string) {
 var wg sync.WaitGroup

 for _, url := range urls {
  wg.Add(1)
  go func(url string) {
   returnType(url)
   wg.Done()
  }(url)
 }

 wg.Wait()
}
```

### Channels

receive <- [channel] <- send

```go
package main

import (
 "fmt"
 "time"
)

func main() {
 ch := make(chan int)
 go func() {
  ch <- 1
 }()

 val := <-ch
 fmt.Printf("Received value: %d\n", val)

 const count = 3
 for i := range count {
  go func() {
   fmt.Println("sending", i)
   ch <- i
  }()
 }

 for range count {
  val := <-ch
  fmt.Println("received", val)
 }

 // close channel when done
 go func() {
  for i := range count {
   fmt.Println("sending", i)
   ch <- i
   time.Sleep(time.Second)
  }
  close(ch)
 }()

 for i := range ch {
  fmt.Println("receiver", i)
 }

 // buffered channel
 ch2 := make(chan int, 1)
 ch2 <- 19
 val2 := <-ch
 fmt.Println(val2)
}
```

### Challenge: channels

```go
package main

import (
 "fmt"
 "net/http"
)

func main() {
 urls := []string{
  "https://www.google.com",
  "https://www.facebook.com",
  "https://www.twitter.com",
  "https://www.mariolazzari.it"}

 ch := make(chan string)
 for _, url := range urls {
  go returnType(url, ch)
 }

 for range urls {
  out := <-ch
  fmt.Println(out)
 }

}

func returnType(url string, out chan string) {
 res, err := http.Get(url)
 if err != nil {
  out <- fmt.Sprintf("%s -> error: %s", url, err)
  return
 }
 defer res.Body.Close()

 ctype := res.Header.Get("Content-Type")
 out <- fmt.Sprintf("%s -> %s", url, ctype)
}
```

### Select

```go
package main

import (
 "fmt"
 "time"
)

func main() {

 ch1, ch2 := make(chan int), make(chan int)

 go func() {
  ch1 <- 41
 }()

 select {
 case val := <-ch1:
  fmt.Printf("Got %d from ch1\n", val)
 case val := <-ch2:
  fmt.Printf("Got %d from ch2\n", val)
 }

 fmt.Println("---")
 out := make(chan float64)
 go func() {
  time.Sleep(100 * time.Millisecond)
  out <- 3.14
 }()

 select {
 case val := <-out:
  fmt.Printf("got %f\n", val)
 case <-time.After(200 * time.Millisecond):
  fmt.Println("Timeout")
 }
}
```

### Context

```go
package main

import (
 "context"
 "fmt"
 "time"
)

type Bid struct {
 AdURL string
 Price float64
}

func main() {
 ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
 defer cancel()
 url := "url"
 bid := findBid(ctx, url)
 fmt.Println(bid)
}

func bestBid(url string) Bid {
 time.Sleep(20 * time.Millisecond)

 return Bid{
  AdURL: url,
  Price: .05,
 }
}

var defaultBid = Bid{
 AdURL: "default",
 Price: .02,
}

func findBid(ctx context.Context, url string) Bid {
 ch := make(chan Bid, 1)
 go func() {
  ch <- bestBid(url)
 }()

 select {
 case bid := <-ch:
  return bid
 case <-ctx.Done():
  return defaultBid
 }
}
```

### Challenge: download size

```go
package main

import (
 "fmt"
 "log"
 "net/http"
 "strconv"
 "time"
)

var (
 urlTemplate = "https://s3.amazonaws.com/nyc-tlc/trip+data/%s_tripdata_2020-%02d.csv"
 colors      = []string{"green", "yellow"}
)

func downloadSize(url string) (int, error) {
 req, err := http.NewRequest(http.MethodHead, url, nil)
 if err != nil {
  return 0, err
 }

 resp, err := http.DefaultClient.Do(req)
 if err != nil {
  return 0, err
 }

 if resp.StatusCode != http.StatusOK {
  return 0, fmt.Errorf(resp.Status)
 }

 return strconv.Atoi(resp.Header.Get("Content-Length"))
}

type result struct {
 url  string
 size int
 err  error
}

func sizeWorker(url string, ch chan result) {
 fmt.Println(url)
 res := result{url: url}
 res.size, res.err = downloadSize(url)
 ch <- res
}

func main() {
 start := time.Now()
 size := 0
 ch := make(chan result)
 for month := 1; month <= 12; month++ {
  for _, color := range colors {
   url := fmt.Sprintf(urlTemplate, color, month)
   go sizeWorker(url, ch)
  }
 }

 for i := 0; i < len(colors)*12; i++ {
  r := <-ch
  if r.err != nil {
   log.Fatal(r.err)
  }
  size += r.size
 }

 duration := time.Since(start)
 fmt.Println(size, duration)
}
```

## Project manager

### Requirements

```sh
go mod init github.com/mariolazzari/go-essential
go tidy
```

### Testing

```sh
go test
```

```go
package sqrt

import (
 "fmt"
 "testing"
)

func almostEqual(v1, v2 float64) bool {
 return Abs(v1-v2) <= 0.001
}

func TestSimple(t *testing.T) {
 val, err := Sqrt(2)

 if err != nil {
  t.Fatalf("error in calculation - %s", err)
 }

 if !almostEqual(val, 1.414214) {
  t.Fatalf("bad value - %f", val)
 }
}

type testCase struct {
 value    float64
 expected float64
}

func TestMany(t *testing.T) {
 testCases := []testCase{
  {0.0, 0.0},
  {2.0, 1.414214},
  {9.0, 3.0},
 }

 for _, tc := range testCases {
  t.Run(fmt.Sprintf("%f", tc.value), func(t *testing.T) {
   out, err := Sqrt(tc.value)
   if err != nil {
    t.Fatal("error")
   }

   if !almostEqual(out, tc.expected) {
    t.Fatalf("%f != %f", out, tc.expected)
   }
  })
 }
}
```

### Testify

[Documentation](https://pkg.go.dev/github.com/jfjallid/testify)

```go
package main

import (
 "testing"

 "github.com/stretchr/testify/require"
)

func TestSimple(t *testing.T) {
 val, err := Sqrt(2)
 require.NoError(t, err)
 require.InDelta(t, 1.414214, val, 0.001)
}
```

### Challenge testing

```go
package sqrt

import (
 "encoding/csv"
 "fmt"
 "io"
 "os"
 "strconv"
 "testing"

 "github.com/stretchr/testify/require"
)

type testCase struct {
 value    float64
 expected float64
}

func loadSqrtCases(t *testing.T) []testCase {
 file, err := os.Open("sqrt_cases.csv")
 require.NoError(t, err)
 defer file.Close()

 var cases []testCase
 rdr := csv.NewReader(file)
 for {
  record, err := rdr.Read()
  if err == io.EOF {
   break
  }
  require.NoError(t, err)

  val, err := strconv.ParseFloat(record[0], 64)
  require.NoError(t, err)
  expected, err := strconv.ParseFloat(record[1], 64)
  require.NoError(t, err)
  tc := testCase{val, expected}
  cases = append(cases, tc)
 }

 return cases
}

func TestMany(t *testing.T) {
 for _, tc := range loadSqrtCases(t) {
  t.Run(fmt.Sprintf("%f", tc.value), func(t *testing.T) {
   out, err := Sqrt(tc.value)
   require.NoError(t, err)
   require.InDelta(t, tc.expected, out, 0.001)
  })
 }
}
```

### Benchmarking and profiling

```go
package main

import (
 "testing"

 "github.com/stretchr/testify/require"
)

var benchText = "Don't communicate by sharing memory, share memory by communicating."

func BenchmarkTokenize(b *testing.B) {
 for i := 0; i < b.N; i++ {
  tokens := Tokenize(benchText)
  require.Equal(b, 10, len(tokens))
 }
}

func TestTokenize(t *testing.T) {
 text := "Who's on first?"
 expected := []string{"who", "s", "on", "first"}
 tokens := Tokenize(text)
 require.Equal(t, expected, tokens)
}
```

## Networking

### JSON

```go

```

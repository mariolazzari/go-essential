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
    var x int
    var y int

    x = 1
    y = 2

    fmt.Printf("x = %v, of type %T\n", x, x)
    fmt.Printf("y = %v, of type %T\n", y, y)

    // type inference :=
    mean := (float64(x) + float64(y)) / 2.0
    fmt.Printf("The mean of %v and %v is: %v\n", x, y, mean)
}
```

### 
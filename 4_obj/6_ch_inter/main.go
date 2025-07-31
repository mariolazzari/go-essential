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

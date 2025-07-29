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

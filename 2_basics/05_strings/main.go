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

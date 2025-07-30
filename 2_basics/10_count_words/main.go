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

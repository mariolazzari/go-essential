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

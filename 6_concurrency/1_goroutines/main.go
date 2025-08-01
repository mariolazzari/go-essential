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

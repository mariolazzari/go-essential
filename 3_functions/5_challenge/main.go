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

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Job struct {
	User   string `json:"user"`
	Action string `json:"action"`
	Count  int    `json:"count"`
}

func main() {
	// GET request
	res, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("GET error on httpbin")
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)

	fmt.Println("----")

	// POST request
	job := &Job{
		User:   "Mario",
		Action: "dev",
		Count:  1,
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(job); err != nil {
		log.Fatalf("Cannot encode job: %s", err)
	}

	res, err = http.Post("https://httpbin.org/post", "application/json", &buf)
	if err != nil {
		log.Fatalf("POST error on httpbin")
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}

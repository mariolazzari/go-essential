package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://httpbin.org/ip", nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	const mb = 1 << 20
	r := io.LimitReader(res.Body, mb)
	io.Copy(os.Stdout, r)
}

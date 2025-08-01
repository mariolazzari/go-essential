package main

import (
	"context"
	"fmt"
	"time"
)

type Bid struct {
	AdURL string
	Price float64
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	url := "url"
	bid := findBid(ctx, url)
	fmt.Println(bid)
}

func bestBid(url string) Bid {
	time.Sleep(20 * time.Millisecond)

	return Bid{
		AdURL: url,
		Price: .05,
	}
}

var defaultBid = Bid{
	AdURL: "default",
	Price: .02,
}

func findBid(ctx context.Context, url string) Bid {
	ch := make(chan Bid, 1)
	go func() {
		ch <- bestBid(url)
	}()

	select {
	case bid := <-ch:
		return bid
	case <-ctx.Done():
		return defaultBid
	}
}

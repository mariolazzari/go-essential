package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()

	val := <-ch
	fmt.Printf("Received value: %d\n", val)

	const count = 3
	for i := range count {
		go func() {
			fmt.Println("sending", i)
			ch <- i
		}()
	}

	for range count {
		val := <-ch
		fmt.Println("received", val)
	}

	// close channel when done
	go func() {
		for i := range count {
			fmt.Println("sending", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Println("receiver", i)
	}

	// buffered channel
	ch2 := make(chan int, 1)
	ch2 <- 19
	val2 := <-ch
	fmt.Println(val2)

}

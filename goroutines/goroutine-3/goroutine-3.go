package main

import (
	"fmt"
	"time"
)

func main() {
	// 1) Create a channel
	ch := make(chan int)

	// 2) Start a goroutine
	go func() {
		fmt.Println("Worker: started work")
		time.Sleep(2 * time.Second) // simulate slow work
		fmt.Println("Worker: finished work")

		// 3) Send result to channel
		ch <- 42
	}()

	// 4) Receive result (blocks until value arrives)
	result := <-ch
	fmt.Println("Main: got result =", result)
}

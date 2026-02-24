package main

import (
	"fmt"
	"time"
)

func main() {
	// Buffered channel allows small queue
	ch := make(chan int, 2)

	// Producer
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("Producing", i)
			ch <- i // blocks only when buffer is full
		}
	}()

	// Consumer (slow)
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Consuming", <-ch)
	}
}

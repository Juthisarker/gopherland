package main

import "fmt"

func main() {
	// Buffered channel with capacity 3
	ch := make(chan int, 3)

	// Send values (does NOT block until buffer is full)
	ch <- 1
	ch <- 2
	ch <- 3

	// Receive values
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

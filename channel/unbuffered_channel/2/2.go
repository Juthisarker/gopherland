package main
// Executable Go program

import (
	"fmt"
	"time"
)

func main() {
	// Create an UNBUFFERED channel
	// No storage inside → send and receive must meet
	ch := make(chan string)

	// Start a worker goroutine
	go func() {
		// Wait to receive a message from main
		// This BLOCKS until main sends
		msg := <-ch
		fmt.Println("Worker received:", msg)

		// Simulate work
		time.Sleep(time.Second)

		// Send response back to main
		// This BLOCKS until main is ready to receive
		ch <- "work done"
	}()

	// Send a message to the worker
	// This BLOCKS until worker receives
	fmt.Println("Main: sending work")
	ch <- "start work"

	// Wait for worker's response
	// This BLOCKS until worker sends
	fmt.Println("Main: waiting for response")
	result := <-ch

	// Print the result
	fmt.Println("Main received:", result)
}

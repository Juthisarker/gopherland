package main

import "fmt"

func sendData(ch chan string) {
	ch <- "First message"
	ch <- "Second message"
	ch <- "Third message"
	close(ch) // Close the channel when done sending
}

func main() {
	ch := make(chan string, 5)

	go sendData(ch)

	// Receive data until the channel is closed
	for msg := range ch {
		fmt.Println("Received:", msg)
	}
}

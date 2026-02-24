package main
// Declares this file as an executable Go program

import (
	"fmt" // Used for printing output
)

func main() {
	// Create an UNBUFFERED channel of type int
	// Unbuffered means: no internal storage
	// Sender and receiver must meet at the same time
	ch := make(chan int)

	// Start a sender goroutine
	go func() {
		// This line executes immediately when the goroutine starts
		fmt.Println("Sender: sending 10")

		// Send value 10 into the channel
		// This BLOCKS until another goroutine is ready to receive
		ch <- 10

		// This line executes ONLY AFTER the receiver has received the value
		fmt.Println("Sender: sent 10")
	}()

	// Main goroutine reaches this line first
	fmt.Println("Receiver: waiting")

	// Receive a value from the channel
	// This BLOCKS until the sender sends a value
	val := <-ch

	// Print the received value
	fmt.Println("Receiver: got", val)
}

// Main creates an unbuffered channel

// Sender goroutine starts

// Sender prints: Sender: sending 10

// Sender reaches ch <- 10 → blocks

// Main prints: Receiver: waiting

// Main reaches <-ch → sender + receiver handshake

// Value 10 is transferred

// Sender unblocks and prints: Sender: sent 10

// Main prints: Receiver: got 10
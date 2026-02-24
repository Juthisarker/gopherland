package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Buffered Channel Example (Email Analogy) ===")
	
	// Like email - you send, they read when they can
	email := make(chan string, 3)  // Buffered channel with capacity 3

	// Sender goroutine
	go func() {  
		fmt.Println("Sender: Sending 3 emails...")
		email <- "Email 1: Meeting notes"
		fmt.Println("  Sent: Email 1")
		email <- "Email 2: Report"
		fmt.Println("  Sent: Email 2")
		email <- "Email 3: Reminder"
		fmt.Println("  Sent: Email 3")
		fmt.Println("Sender: All emails sent immediately! (queued in inbox)")
	}()

	fmt.Println("Main: Receiver is busy for 2 seconds...")
	time.Sleep(2 * time.Second)  // Simulate receiver being busy

	fmt.Println("\nMain: Receiver checking email now...")
	// Receiver checks email later
	for i := 0; i < 3; i++ {
		msg := <-email
		fmt.Printf("Receiver: Read %d: %s\n", i+1, msg)
		time.Sleep(500 * time.Millisecond) // Simulate reading time
	}

	fmt.Println("\n=== All emails processed! ===")
}
package main

// Declares this file as an executable Go program

import (
	"fmt"  // Provides functions for formatted output (printing)
	"sync" // Provides synchronization primitives like WaitGroup
)

func main() {
	// Declare a WaitGroup to wait for multiple goroutines to finish
	var wg sync.WaitGroup

	// Increase the WaitGroup counter to 2
	// This tells the WaitGroup that we are waiting for two goroutines
	wg.Add(2)

	// Start the first goroutine (Task 1)
	go func() {
		// Ensure Done() is called when this goroutine finishes
		// This decreases the WaitGroup counter by 1
		defer wg.Done()

		// Print that Task 1 has completed
		fmt.Println("Task 1 done")
	}()

	// Start the second goroutine (Task 2)
	go func() {
		// Ensure Done() is called when this goroutine finishes
		// This decreases the WaitGroup counter by 1
		defer wg.Done()

		// Print that Task 2 has completed
		fmt.Println("Task 2 done")
	}()

	// Block the main goroutine until the WaitGroup counter becomes 0
	// This happens after both goroutines call wg.Done()
	wg.Wait()

	// This line runs only after both goroutines have finished
	fmt.Println("Both tasks finished")
}

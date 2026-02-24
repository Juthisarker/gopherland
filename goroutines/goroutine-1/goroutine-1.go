package main
// Declares this file as an executable Go program

import (
	"fmt"   // Used for printing output to the console
	"sync"  // Provides synchronization primitives (WaitGroup)
	"time"  // Used for sleeping to simulate work
)

// worker represents a unit of work that runs in its own goroutine
func worker(id int, wg *sync.WaitGroup) {
	// Ensure that wg.Done() is called when this function finishes
	// This decrements the WaitGroup counter by 1
	defer wg.Done()

	// Print that this worker has started working
	fmt.Printf("Worker %d is working...\n", id)

	// Simulate some work by sleeping for 1 second
	// This blocks ONLY this goroutine, not the whole program
	time.Sleep(time.Second)

	// Print that this worker has finished its work
	fmt.Printf("Worker %d finished!\n", id)
}

func main() {
	// Declare a WaitGroup to track active goroutines
	var wg sync.WaitGroup

	// Start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		// Increase the WaitGroup counter by 1
		// This tells main(): "one more goroutine will finish later"
		wg.Add(1)

		// Start a new goroutine running the worker function
		// The goroutine receives:
		//  - a unique worker id
		//  - a pointer to the shared WaitGroup
		go worker(i, &wg)
	}

	// Block the main goroutine until the WaitGroup counter becomes 0
	// This happens after all workers call wg.Done()
	wg.Wait()

	// This line runs only AFTER all worker goroutines finish
	fmt.Println("All workers done!")
}

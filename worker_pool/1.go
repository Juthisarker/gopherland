package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a channel to send jobs (integers) to workers.
	// This is an unbuffered channel, so sends will block until a worker receives.
	jobs := make(chan int)

	// WaitGroup is used to wait for all workers to finish.
	var wg sync.WaitGroup

	// How many workers we want to run in parallel.
	workerCount := 3

	// Start a fixed number of workers.
	for i := 0; i < workerCount; i++ {
		// Each worker is one goroutine, so we say "one more thing to wait for".
		wg.Add(1)

		// Launch a worker goroutine.
		// We pass 'i' as workerID to avoid closure-capture issues.
		go func(workerID int) {
			// When this worker function returns, mark it as done in the WaitGroup.
			defer wg.Done()

			// Continuously receive jobs from the 'jobs' channel.
			// 'range jobs' loops until the channel is closed.
			for job := range jobs {
				// Simulate processing the job (here: just print).
				fmt.Println("worker", workerID, "processing job", job)
			}

			// When the channel is closed and drained, the loop ends
			// and the worker exits, triggering wg.Done() via defer.
		}(i)
	}

	// Send 5 jobs into the jobs channel.
	// Each value 'j' represents one unit of work.
	for j := 1; j <= 5; j++ {
		jobs <- j // This will block until some worker reads from 'jobs'.
	}

	// No more jobs will be sent, so close the channel.
	// This allows workers' 'range jobs' loops to terminate.
	close(jobs)

	// Wait until all workers have finished processing.
	// They finish when the jobs channel is closed AND all jobs are consumed.
	wg.Wait()

	fmt.Println("all workers finished")
}

package main

import "fmt"

func countToFive(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch) // Close channel after sending all data
}

func main() {
	ch := make(chan int)

	go countToFive(ch) // Start the goroutine

	// Receive the numbers from the channel and print them
	for num := range ch {
		fmt.Println(num)
	}
}

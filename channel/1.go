package main

import "fmt"

// Define the function separately
func sendData(ch chan int) {
    ch <- 42 // Sending data into the channel
}

func main() {
    // Create a channel of type int
    ch := make(chan int)

    // Start a goroutine with a named function
    go sendData(ch)

    // Receive data from the channel
    result := <-ch
    fmt.Println("Received:", result)
}

/*
Summary of Execution Flow:
1. The main function starts and creates a channel.

2. A goroutine is launched using go sendData(ch).

3. The sendData function runs in the goroutine and sends 42 into the channel.

4. The main function waits for data from the channel (<-ch), and the main goroutine is blocked until data is received.

5. The goroutine sends the value 42 to the channel.

6. The main goroutine receives the value 42 from the channel.

7. The result is printed (Received: 42).

8. The program ends.

 **Both the sending and receiving goroutines block on the channel. The sending goroutine blocks until the data is received, and the main goroutine blocks until the data is sent.

*/
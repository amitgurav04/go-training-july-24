package main

import (
	"fmt"
)

func main() {
	// Create a buffered channel with a capacity of 3
	ch := make(chan int, 3)
	// Sender goroutine
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- i // Send values to the channel
			fmt.Println("Sending", i)
		}
	}()
	// Receiver goroutine
	for value := range ch {
		fmt.Println("Received", value)
	}
}

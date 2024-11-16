package main

import (
	"fmt"
	"sync"
)

func generateNumbers(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Notify the WaitGroup when the goroutine is done
	for i := 1; i <= 5; i++ {
		ch <- i // Send data to the channel
	}
}

func printNumbers(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Notify the WaitGroup when the goroutine is done
	for i := 1; i <= 5; i++ {
		num := <-ch // Receive data from the channel
		fmt.Println("Received:", num)
	}
}

func main() {
	ch := make(chan int)   // Create a channel
	wg := sync.WaitGroup{} // Initialize the WaitGroup

	wg.Add(2) // We are going to wait for 2 goroutines

	go generateNumbers(ch, &wg) // Start the goroutine to generate numbers
	go printNumbers(ch, &wg)    // Start the goroutine to print numbers

	wg.Wait() // Wait for both goroutines to finish
	fmt.Println("All numbers received and printed.")
}

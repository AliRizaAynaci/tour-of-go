package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Called when the goroutine is done
	fmt.Printf("Task %d started\n", id)
}

func main() {
	wg := sync.WaitGroup{} // Creates a WaitGroup to track pending goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1)       // Increases the number of goroutines to wait for
		go task(i, &wg) // Starts each goroutine
	}

	wg.Wait()                          // Waits for all goroutines to finish
	fmt.Println("All tasks completed") // Prints a message once all goroutines are done
}

package main

import (
	"fmt"
	"sync"
)

var counter int
var lock sync.Mutex

func increment() {
	lock.Lock()   // Acquire the lock
	counter++     // Update the shared resource
	lock.Unlock() // Release the lock
}

func main() {

	wg := sync.WaitGroup{} // Create a WaitGroup to wait for goroutines to finish

	for i := 0; i < 1000; i++ {
		wg.Add(1) // Increase the number of goroutines to wait for
		go func() {
			defer wg.Done() // Decrease the counter when the goroutine finishes
			increment()     // Increment the counter
		}()
	}

	wg.Wait()                         // Wait until all goroutines are finished
	fmt.Println("Counter: ", counter) // Print the final value of the counter
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	// waitgroup for synchronization
	wg := sync.WaitGroup{}
	// wring number of goroutine to wait
	wg.Add(2)

	go func() {
		fmt.Println("Hello from go routine")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Hello concurrency")
}

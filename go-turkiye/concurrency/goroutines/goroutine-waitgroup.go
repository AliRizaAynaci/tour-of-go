package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		fmt.Println("Hello from goroutine")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Hello from main")

}

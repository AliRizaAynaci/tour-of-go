package main

import "fmt"

func main() {

	unbufferedChan := make(chan int)

	// reader goroutine
	go func(unbuffChan chan int) {
		value := <-unbuffChan
		fmt.Println(value)
	}(unbufferedChan)

	// writer goroutine
	go func(unbuffChan chan int) {
		unbuffChan <- 1
	}(unbufferedChan)

	fmt.Println("Hello channels")

	/*
		Output is non-deterministic. Scheduler probably will not have time to schedule goroutines.
		So we will not see channel value in the output
	*/
}

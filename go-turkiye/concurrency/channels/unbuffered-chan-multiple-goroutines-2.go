package main

import (
	"fmt"
	"time"
)

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

	// now scheduler have time to schedule, still, this might not work and not the best solution
	time.Sleep(time.Second)

	fmt.Println("Hello channels")
}

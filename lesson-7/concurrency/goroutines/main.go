package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {

	// "go say("world")" starts a new goroutine, is a lightweight thread managed by the Go runtime.
	// This allows the "say" function to run with the "world" parameter as a separate goroutine (a lightweight thread).
	// This way, other tasks in the main function can continue running concurrently.
	go say("world")

	// The "say("hello")" call, however, runs directly in the main goroutine.
	// This call blocks the main function until it completes, meaning the "main" function will not finish until "say("hello")" is done.
	say("hello")

}

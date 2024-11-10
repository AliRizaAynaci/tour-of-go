package main

import "fmt"

func main() {

	go func() {
		fmt.Println("Hello from goroutine")
	}()

	fmt.Println("Hello from main")
}

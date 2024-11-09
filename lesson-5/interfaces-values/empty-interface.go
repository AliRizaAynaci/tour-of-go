package main

import "fmt"

func main() {

	var i interface{}
	describe2(i) // return <nil>, <nil>

	i = 42
	describe2(i) // return 42, int

	i = "Hello"
	describe2(i) // return Hello, string

}

func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

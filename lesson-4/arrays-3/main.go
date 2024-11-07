package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

	var s []int

	// func append(s []T, vs ...T) []T
	// append works on nil slices
	s = append(s, 0)
	printSlice(s)

	// the slice is grows as needed
	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)

	// i is index of pow array, v is value of pow array
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow1 := make([]int, 10)
	for i := range pow1 {
		pow1[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow1 {
		fmt.Printf("%d\n", value)
	}
}

func printSlice(s []int) {
	fmt.Printf("len = %d cap = %d arr = %v\n", len(s), cap(s), s)
}

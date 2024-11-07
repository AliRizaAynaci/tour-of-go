package main

import (
	"fmt"
	"strings"
)

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}

	//Creating a slice with make
	//Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
	//
	//	The make function allocates a zeroed array and returns a slice that refers to that array:

	a := make([]int, 5) // its meaning is len = 5, cap = 5, also it fills in the array with 0 through len
	printSlice("a", a)

	b := make([]int, 0, 5) // its meaning is len = 0, cap = 5
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

}

func printSlice(s string, arr []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(arr), cap(arr), arr)
}

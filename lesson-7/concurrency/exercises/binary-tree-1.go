package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)  // Traverse left subtree
	ch <- t.Value     // Send the root value to the channel
	Walk(t.Right, ch) // Traverse right subtree
}

// Same determines whether the trees t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	// Start Walk in goroutines for both trees
	go func() {
		Walk(t1, ch1)
		close(ch1) // Close the channel once done walking
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2) // Close the channel once done walking
	}()

	// Compare values received from both channels
	for v1 := range ch1 {
		v2, ok := <-ch2
		if !ok || v1 != v2 { // If values differ or channel is closed, trees are different
			return false
		}
	}
	_, ok := <-ch2 // Ensure both channels are exhausted
	return !ok
}

func main() {
	// Create a new channel
	ch := make(chan int)

	// Start the Walk function in a goroutine
	go Walk(tree.New(1), ch)

	// Read and print 10 values from the channel
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch) // Print the value received from the channel
	}

	// Test the Same function
	fmt.Println(Same(tree.New(1), tree.New(1))) // Should return true
	fmt.Println(Same(tree.New(1), tree.New(2))) // Should return false
}

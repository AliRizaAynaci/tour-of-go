package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Append(val T) {
	newNode := &List[T]{val: val}
	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (l *List[T]) Print() {
	current := l
	for current != nil {
		fmt.Printf("%v ->", current.val)
		current = current.next
	}
}

func main() {
	// Create a new list with an initial value
	head := &List[int]{val: 1}

	// Append elements to the list
	head.Append(2)
	head.Append(3)
	head.Append(4)

	// Print the list elements
	head.Print()
}

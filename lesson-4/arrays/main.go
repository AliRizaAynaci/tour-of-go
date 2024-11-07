package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// SLICES
	var ss []int = primes[1:4]
	fmt.Println(ss)

	// Slices are like references to arrays
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	x := names[0:2]
	y := names[1:3]
	fmt.Println(x, y)

	y[0] = "XXX"
	fmt.Println(x, y)
	fmt.Println(names)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	newArray := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(newArray)

	newArray = newArray[1:4]
	fmt.Println(newArray)

	newArray = newArray[:2]
	fmt.Println(newArray)

	newArray = newArray[1:]
	fmt.Println(newArray)

	arr := []int{2, 3, 5, 7, 11, 13}
	printSlice(arr)

	// Slice the slice to give it zero length.
	arr = arr[:0]
	printSlice(arr)

	// Extend its length.
	arr = arr[:4]
	printSlice(arr)

	// Drop its first two values.
	arr = arr[2:]
	printSlice(arr)

}

func printSlice(arr []int) {
	fmt.Printf("len = %d cap = %d arr = %v\n", len(arr), cap(arr), arr)
}

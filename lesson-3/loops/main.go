package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	x := 1

	for x < 1000 {
		x += x
	}

	fmt.Println(x)

	y := 1
	for y < 100 {
		y += y
	}
	fmt.Println(y)

	for {
		// infinitive loop
	}
}

package main

import "fmt"

type I1 interface {
	M()
}

func main() {
	var i I1
	describe1w(i)
	i.M()
}

func describe1(i I1) {
	fmt.Printf("(%v, %T)\n", i, i)
}

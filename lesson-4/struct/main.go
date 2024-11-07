package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{3, 4}
	v.X = 5
	fmt.Println(v.X)
	p := &v
	p.X = 1e9
	fmt.Println(v)

	fmt.Println("-------------")
	fmt.Println(v1, p, v2, v3)
}

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//func (v Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
//func (v *Vertex) Scale(f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}

//func Abs(v Vertex) float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
//func Scale(v *Vertex, f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}

//func (v *Vertex) Scale(f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}
//
//func ScaleFunc(v *Vertex, f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// v := Vertex{3, 4}
	// v.Scale(10)
	// fmt.Println(v.Abs())

	//v := Vertex{3, 4}
	//Scale(&v, 10)
	//fmt.Println(Abs(v))

	//v := Vertex{3, 4}
	//v.Scale(2)
	//ScaleFunc(&v, 10)
	//
	//p := &Vertex{4, 3}
	//p.Scale(3)
	//ScaleFunc(p, 8)
	//
	//fmt.Println(v, p)

	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}

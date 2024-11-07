package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

var m1 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
	fmt.Println(m1)

	mp := make(map[string]int)

	mp["Answer"] = 42
	fmt.Println("The value:,", mp["Answer"])

	mp["Answer"] = 48
	fmt.Println("The value:", mp["Answer"])

	delete(mp, "Answer")
	fmt.Println("The value:", mp["Answer"])

	v, ok := mp["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

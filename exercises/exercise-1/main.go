package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0           // Starting guess
	threshold := 1e-10 // Small threshold for stopping condition
	for i := 0; i < 10; i++ {
		zNew := z - (z*z-x)/(2*z) // Apply Newton's method formula
		fmt.Printf("Iteration %d: z = %v\n", i+1, z)
		if math.Abs(z-zNew) < threshold {
			return zNew
		}
		z = zNew
	}
	return z
}

func main() {
	fmt.Println("Computed Sqrt(2):", Sqrt(2))
	fmt.Println("Actual math.Sqrt(2):", math.Sqrt(2))
}

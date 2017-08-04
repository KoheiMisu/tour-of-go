package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64, tryCount int) float64 {
	for i := 1; i <= tryCount; i++ {
		nextVal := float64(x - (x*x-2)/(x*2))
		if math.Abs(nextVal-x) < 0.0000000000000001 {
			fmt.Printf("end calcurate at %v\n", i)
			break
		}

		x = nextVal
	}

	return x
}

func main() {
	fmt.Printf("math.Sqrt() => %v\n\n", math.Sqrt(2))
	fmt.Printf("My Sqrt() => %v\n", Sqrt(2, 3))
	fmt.Printf("My Sqrt() => %v\n", Sqrt(2, 10))
}

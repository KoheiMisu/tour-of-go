package main

import (
	"fmt"
	"math"
)

func main() {
	q := []int{2, 3, 5, 7, 11, 13}

	for i, val := range q {
		fmt.Printf("index: %d, pow: %f\n", i, math.Pow(float64(val), 2))
	}
}

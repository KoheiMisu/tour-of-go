package main

import (
	"fmt"
	"math"
)

func powElse(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf(""+
			"[Over Limit..] %g >= %g\n", v, lim)
	}

	return lim
}

func main() {
	fmt.Println(
		powElse(3, 2, 10), // => 9
		powElse(3, 3, 10), // => 10
	)
}

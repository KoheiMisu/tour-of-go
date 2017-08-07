package main

import (
	"math"
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("invalid value at %f", e)
}

func Sqrt(x float64, tryCount int) (float64, error) {
	if x < 0 {
		// typeに引数が渡せる...
		return 0, ErrNegativeSqrt(x)
	}

	for i := 1; i <= tryCount; i++ {
		nextVal := float64(x - (x*x-2)/(x*2))
		if math.Abs(nextVal-x) < 0.0000000000000001 {
			fmt.Printf("end calcurate at %v\n", i)
			break
		}

		x = nextVal
	}

	return x, nil
}

func main() {
	if val, err := Sqrt(-2, 5); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("My Sqrt() => %v\n", val)
	}
}

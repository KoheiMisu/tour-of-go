package main

import (
	"fmt"
	"math"
)

// 内部で使う値は決まっているが、ロジックを入れ替えたい場合などに
// 関数を引数として受けて使うとよさそう
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	// 変数v, 型Tにおいて、T(v)は変数vをT型へと変換する
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	// z := uint(f)でももちろん可能

	fmt.Println(x, 4, z)
}

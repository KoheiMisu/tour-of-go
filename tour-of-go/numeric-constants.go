package main

import "fmt"

// 数値の定数は高精度な値(Value)
// 型のない定数は、その状況によって必要な型を取ることになる
// int は64bitの整数を保持できるがそれでは足りないときなどに使う
const (
	Big   = 1 << 100
	Small = Big >> 99 // = 1 << 1
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	//fmt.Println(Big)
	//fmt.Println(Small)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1            // shift operator. シフト演算とは、2進数のビットパターンを右または左にずらす演算である
	z      complex128 = cmplx.Sqrt(-5 + 12i) // default (0+0i)
)

func main() {
	const f = "%T(%v)\n"

	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}

//1 << 1 = 2
//1 << 2 = 4
//1 << 3 = 8
//1 << 4 = 16
//1 << 5 = 32
//1 << 6 = 64
//1 << 7 = 128
//1 << 8 = 256
//1 << 9 = 512

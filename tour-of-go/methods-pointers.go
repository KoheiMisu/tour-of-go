package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ポインタレシーバを持つメソッドはレシーバが指す変数を変更できる
// レシーバ自信を更新することが多いので、変数レシーバよりもポインタレシーバの方が一般的
// 変数レシーバでは、元のstructをコピーして操作を行う。そのため、元のstructは変更されない
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

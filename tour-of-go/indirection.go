package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 一般的には変数レシーバとポインタレシーバのどちらかですべてのメソッドを揃える
// 理由
//

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func StaticScaleFunc(v Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)
	fmt.Println(v, p)

	// 関数にstructを渡す場合は、型が合っていないと正しく動かない
	//ScaleFunc(v) => compile error

	// 変数レシーバの場合は、変数でもポインタでもレシーバとして取ることが可能
	// p.StaticScaleFuncは内部的に(*p).StaticScaleFunc()として解釈される
}

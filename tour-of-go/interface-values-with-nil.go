package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

// interface自体の中にある具体的な値がにnilの場合、メソッドはnilをレシーバーとして呼ばれる

// Go ではnilをレシーバーとして呼び出されても適切に処理するメソッドを記述するのが一般的
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}

	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T

	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)", i, i)
}

package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i) // => (<nil>, <nil>)

	// 下記の処理ではエラーになる
	// nilインターフェースの値は、値も具体的な型も保持しない
	// そのため、具体的なメソッドを示す型がインターフェースのタプル内に存在しない
	// 実行するとランタイムエラーになる
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)", i, i)
}

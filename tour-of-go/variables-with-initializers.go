package main

import "fmt"

var i, j int = 1, 2

func main() {
	// 初期化の値が渡されているときは、型を省略することができる
	// 変数は初期化値が持つ型になる
	var c, python, java = true, false, "no!"

	fmt.Println(i, j, c, python, java) // => 1 2 true false no!
}

package main

import "fmt"

func main() {
	// 関数自体はすぐに評価されるが、関数自体は
	// 呼び出し元の関数がreturnされるまで実行されない
	defer fmt.Println(" deferred world")

	fmt.Println("Hello")
}

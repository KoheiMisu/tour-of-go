package main

import "fmt"

func main() {
	fmt.Println("counting")

	// deferに複数の関数を指定するとスタックが行われる
	// 呼び出し順はLIFOなので一番最近入れられたものから
	// 関数のcallが行われる
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

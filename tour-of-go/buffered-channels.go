package main

import "fmt"

func main() {
	// バッファ2のチャンネルを作成
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 バッファ以上のチャンネルには書き込めない. deadlock

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//fmt.Println(<-ch) 全てのチャンネルを受信した後に読み込むとエラー deadlockになる
}

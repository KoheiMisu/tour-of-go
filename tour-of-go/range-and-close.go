package main

import "fmt"

func fibonacci(n int, c chan int) {
	fmt.Println(n)
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	// 全ての値を受信するためのループ文
	// go func()だけではすぐに処理が終了するのでチャンネルに送信される値を
	// 拾ってやる必要がある
	for i := range c {
		fmt.Println(i)
	}
}

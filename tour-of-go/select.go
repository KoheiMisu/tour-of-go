package main

import "fmt"

// selectは複数あるcaseのいずれかが準備できるようになるまでブロックし
// 準備ができたcaseを実行する

// 複数のcaseの準備ができている場合、caseはランダムに選択される

func fibonacci(c chan int, q chan int) {
	x, y := 0, 1

	for {
		select {
		// 送信ができるようになるまで待機
		case c <- x:
			x, y = y, x+y
		// 受信ができるようになるまで待機
		case <-q:
			fmt.Println("end")
			return

		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			// channelの受信は同期なので送信されるまで
			// ここで待つことができる

			// そもそもgo func()は非同期なのでいくらここで止まってるからといって
			// メインプロセスに影響はない
			fmt.Println(<-c)

			// for 節の分、受信したら終了する
		}

		quit <- 0
	}()

	fibonacci(c, quit)
}

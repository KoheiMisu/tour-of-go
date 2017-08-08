package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk関数はtreeからすべての値をtreeからチャンネルに送信する
func Walk(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}
		// leftのtreeで探索を繰り返す
		walker(t.Left)
		// leftの探索が終わったら値を送信する
		ch <- t.Value
		// rightのtreeで探索を繰り返す
		walker(t.Right)
	}
	walker(t)
	close(ch)
}

// Same関数は tree t1とtree t2が同じ値を含むかどうかを決定する
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		// channelがcloseされるとok*にはfalseが入る
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if v1 != v2 || ok1 != ok2 {
			return false
		}

		if !ok1 {
			break
		}
	}

	return true
}

/**
やりたいこと

二つの二分木が同じ順序で値を保持しているかどうかを確認する

*/
func main() {
	fmt.Println("1 and 1 same: ", Same(tree.New(1), tree.New(1)))
	fmt.Println("1 and 2 same: ", Same(tree.New(1), tree.New(2)))
}

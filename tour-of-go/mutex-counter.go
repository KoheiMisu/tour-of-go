package main

import (
	"fmt"
	"sync"
	"time"
)

/**
goroutineがチャンネルとのコミュニケーションを必要としない時に
変数への同時アクセスを防ぐためにはmutexを使う
*/

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// LockとUnLockを使うことで排他制御を実現できる
// 下記ではmapには一つのgoroutineしかアクセスできないようになっている
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

// 複数のgoroutineが同時アクセスすることを防ぐ
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	//time.Sleep(time.Second) 複数のgoroutineがアクセスしないか確認できる
	defer c.mux.Unlock() // Unlockを保障するためにはdeferが一般的かなと。コケてロック解除できなかったら辛い
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
		go c.Inc("test")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
	fmt.Println(c.Value("test"))
}

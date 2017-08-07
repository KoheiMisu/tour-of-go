package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func sumDelay(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	time.Sleep(time.Second)
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sumDelay(s[len(s)/2:], c)

	// cから受信されるまではブロックされる
	// sumDelayで意図的に処理を遅らせているが処理を待ってくれる
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

}

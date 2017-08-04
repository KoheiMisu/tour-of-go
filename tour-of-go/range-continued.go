package main

import "fmt"

func main() {
	pow := make([]int, 10)

	// valueが不要な場合
	for index := range pow {
		pow[index] = 1 << uint(index)
	}

	// index が不要な場合
	for _, val := range pow {
		fmt.Printf("%d\n", val)
	}

}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seedを与えないと常に同じ値を返す
	fmt.Println("My favorite number is", rand.Intn(10)) // => 1

	rand.Seed(time.Now().UnixNano())
	fmt.Println("(Seed Given) My favorite number is", rand.Intn(10)) // => 1 ~ 10
}

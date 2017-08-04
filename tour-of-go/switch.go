package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on")

	// breakさせずに下に通す場合はfallthrough文を使う
	// defaultではマッチしたらbreakされる
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.", os)
	}
}

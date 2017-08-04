package main

import "fmt"

const Pi = 3.14

func main() {

	// constは、character, string, boolean, numericのみで使える

	const World = "世界"

	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go Rules?", Truth)
}

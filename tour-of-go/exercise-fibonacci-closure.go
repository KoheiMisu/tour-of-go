package main

import "fmt"

func fibonacci() func() int {
	var next, prev, tmp int

	return func() int {
		if next == 0 {
			next++
			return 0
		}

		if next == 1 && tmp == 0 {
			tmp++
			return 1
		}

		tmp = next
		next = next + prev
		prev = tmp
		return next
	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

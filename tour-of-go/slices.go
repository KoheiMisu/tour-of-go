package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // x[a:b] => a <= index < b の範囲で切り抜く
	fmt.Println(s)            // => [3, 5, 7, 11]

	// pointer
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// sliceは配列の参照であることを実験
	// 同じ元となる配列を共有しているスライスは、他方で変更があった場合に
	// それが反映される
	b[0] = "XXXXXXXXXX"
	fmt.Println(a, b)  // => [John XXXXXXXXXX] [XXXXXXXXXX George]
	fmt.Println(names) // => [John XXXXXXXXXX George Ringo]

}

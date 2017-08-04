package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	printSlice(q)

	// 長さ0, 空のスライスを作る.(ただ、capは保持されたまま)
	q = q[:1]
	printSlice(q)

	// 長さ0, 空のスライスを作る.(ただ、capは保持されたまま)
	q = q[:0]
	printSlice(q)

	// スライスの拡張
	// というか、中身の復号的なことができる
	q = q[:4]
	printSlice(q)

	// 先頭の2要素を削除する
	// [length:]で指定したときのみ、cap値が変わる
	q = q[2:]
	printSlice(q) // len=2, cap=4, [5 7]

	// sliceのゼロ値はnil(sliceがempty かつ length 0 かつ cap 0)
	var nil_s []int
	fmt.Println(nil_s, len(nil_s), cap(nil_s))

	if nil_s == nil {
		fmt.Println("nil!")
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d, %v\n", len(s), cap(s), s)
}

package main

import "fmt"

func main() {
	i, j := 42, 2701

	// &オペレータで、該当変数のポインタを引き出す
	p := &i

	// *　オペレータは、ポインタが指す先の変数を示す
	fmt.Println(*p)

	// ポインタを通して変数を代入する
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

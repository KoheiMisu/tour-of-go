package main

import (
	"fmt"
	"util"
)

func main() {
	a := make([]int, 5)
	util.PrintSlice("a", a) // length を指定

	b := make([]int, 0, 5) // length 0, cap 5 のslice
	util.PrintSlice("b", b)

	c := b[:2]
	util.PrintSlice("c", c)

	d := c[2:5]
	fmt.Printf("%s len=%d cap=%d %v\n",
		"d", len(d), cap(d), d)
	util.PrintSlice("d", d)
}

package main

import "util"

func main() {
	var s []int
	util.PrintSlice("s_01", s)

	s = append(s, 0)
	util.PrintSlice("s_02", s)

	s = append(s, 1)
	util.PrintSlice("s_03", s)

	// 第二引数以降は、順番にsliceに値が登録されていく
	s = append(s, 2, 3, 4)
	util.PrintSlice("s_04", s)
}

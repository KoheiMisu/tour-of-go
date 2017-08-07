package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The Value: ", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The Value: ", m["Answer"])

	// mapでは要素の削除が行える
	delete(m, "Answer")
	fmt.Println("The Value: ", m["Answer"])

	// 指定したキーが存在するかどうか. phpでいうissetが気軽に行える
	v, ok := m["Answer"]
	fmt.Println("The Value: ", v, ", Present?", ok)

	if _, ok := m["Answer"]; ok {
		fmt.Println("exists")
	}
}

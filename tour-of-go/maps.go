package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// mapのゼロ値はnil, nilマップはキーを持っておらず、またキーを追加することもできない
var m map[string]Vertex

func main() {
	// make関数で指定した型の初期化され利用できるようにしたmapを取得する
	m = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{
		40.68443, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

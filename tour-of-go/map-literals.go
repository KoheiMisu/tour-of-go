package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.0000, -74.0000,
	},
	"Google": Vertex{
		24.0000, -35.0000,
	},
}

// mapに渡すトップレベルの型が単純な型名である場合はリテラルの要素から推定が可能
// 型名は省略ができる
var m2 = map[string]Vertex{
	"Bell Labs": {
		40.0000, -74.0000,
	},
	"Google": {
		24.0000, -35.0000,
	},
}

func main() {
	fmt.Println(m)
	fmt.Println(m2)
}

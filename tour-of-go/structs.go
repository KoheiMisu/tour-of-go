package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	fmt.Printf("%T\n", Vertex{1, 2}) // => main.Vertex

	// access struct field
	v := Vertex{1, 3}
	v.X = 8
	fmt.Println(v.X) // => 8

	// struct pointers
	p := &v
	p.X = 1e9
	fmt.Println(v) // => {1000000000, 3}
}

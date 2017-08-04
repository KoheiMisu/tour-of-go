package main

import "fmt"

var v []interface{}

func main() {
	v = append(v, 1)
	v = append(v, "string")
	v = append(v, false)
	v = append(v, 0.93)
	v = append(v, 'b')
	v = append(v, 0.867+0.5i)

	for _, val := range v {
		fmt.Printf("%v is of type %T\n", val, val)
	}
}

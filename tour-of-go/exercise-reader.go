package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct {
}

func (mr MyReader) Read(b []byte) (n int, e error) {
	for i, _ := range b {
		b[i] = 'A'
		n++
	}
	return
}

// add Read([]byte) (int, error)

func main() {
	reader.Validate(MyReader{})
}

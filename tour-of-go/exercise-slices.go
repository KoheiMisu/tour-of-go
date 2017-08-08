package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	// 長さdyのsliceを作る
	image := make([][]uint8, dy)

	// golangでは image[] = [uint8(1)]のような書き方はできない
	// まずはindex -> yに対応する[]uint8をmakeする
	for y := range image {
		image[y] = make([]uint8, dx)
	}

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			image[y][x] = uint8(x)
		}
	}

	return image
}

func main() {
	pic.Show(Pic)
}

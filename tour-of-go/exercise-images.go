package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
}

func (i *Image) ColorModel() color.Model {
	c := color.RGBAModel
	return c
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 200)
}

func (i *Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	m := new(Image)
	pic.ShowImage(m)

	// 上記のようにポインタのレシーバでは
	// m := Image{} とやるとmはポインタ型ではないのでレシーバとして認識されない
	// m := new(Type) か m := &Image{}でポインタを受け取るようにする
}

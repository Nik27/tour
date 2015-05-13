package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
	"math"
)

type Image struct{}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 500, 500)
}

func (Image) At(x, y int) color.Color {
	x = int(math.Abs(float64(x - 250)))
	y = int(math.Abs(float64(y - 250)))
	x = int(math.Pow(float64(x), 2.0/3.0))
	y = int(math.Pow(float64(y), 2.0/3.0))
	if x+y == int(math.Pow(250.0, 2.0/3.0)) {
		return color.RGBA{255, 255, 255, 255}
	} else if x == 0 || y == 0 {
		return color.RGBA{255, 0, 0, 128}
	}
	return color.RGBA{0, 0, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}

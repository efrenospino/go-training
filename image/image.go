package main

import "golang.org/x/tour/pic"

import "image"
import "image/color"

type EfrensImage struct{}

// ColorModel sfghs hdsguds gd gh ggjs
func (ei EfrensImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (ei EfrensImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (ei EfrensImage) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	//	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	efrensPic := EfrensImage{}
	pic.ShowImage(efrensPic)
}

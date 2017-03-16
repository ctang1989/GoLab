package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

// create the custom image type
type Image struct {
	width, height int
	pColor        uint8
}

// set the image dimensions and location
func (self *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, self.width, self.height)
}

// set the image color model
func (self *Image) ColorModel() color.Model {
	return color.RGBAModel
}

// set the image color per pixel
func (self *Image) At(x, y int) color.Color {
	return color.RGBA{self.pColor + uint8(x), self.pColor + uint8(y), 255, 255}
}

// create and show the image
func main() {
	m := Image{100, 100, 128}
	pic.ShowImage(&m)
}

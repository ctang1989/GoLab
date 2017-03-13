package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	// initialize the image
	image := make([][]uint8, dy)

	for x := 0; x < dy; x++ {
		image[x] = make([]uint8, dx)
		for y := 0; y < dx; y++ {

			// set a color value for each pixel
			image[x][y] = uint8((x ^ y) * (x ^ y))
		}
	}
	return image
}

func main() {
	pic.Show(Pic)
}

// https://nylira.com/a-tour-of-go-solutions/

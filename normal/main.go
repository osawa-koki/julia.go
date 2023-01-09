package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
	iterations             = 200
	contrast               = 15
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, julia(z))
		}
	}
	file, _ := os.Create("julia.png")
	png.Encode(file, img)
	file.Close()
}

func julia(z complex128) color.Color {
	var v complex128 = complex(-0.8, 0.156)
	for n := uint8(0); n < iterations; n++ {
		z = z*z + v
		if cmplx.Abs(z) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

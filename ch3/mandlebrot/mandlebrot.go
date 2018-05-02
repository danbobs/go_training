// Mandlebrot creates a png of a Mandlebrot plot
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, 2, 2
	width, height          = 1024, 1024
	iterations             = 200
	contrast               = 15
)

func main() {

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y1 := float64(py)/height*(ymax-ymin) + ymin
		y2 := (float64(py)+0.5)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x1 := float64(px)/width*(xmax-xmin) + xmin
			x2 := (float64(px)+0.5)/width*(xmax-xmin) + xmin
			z1, z2, z3, z4 := complex(x1, y1), complex(x1, y2), complex(x2, y1), complex(x2, y2)
			m1, m2, m3, m4 := mandlebrot(z1), mandlebrot(z2), mandlebrot(z3), mandlebrot(z4)
			average := (float32(m1) + float32(m2) + float32(m3) + float32(m4)) / 4
			col := colorFromIterations(uint8(average))
			img.Set(px, py, col)
		}
	}
	png.Encode(os.Stdout, img)
}

func mandlebrot(z complex128) uint8 {

	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return n
		}
	}
	return iterations
}

func colorFromIterations(n uint8) color.Color {

	if n == iterations {
		return color.Black
	}

	r := 255 - n*contrast/3
	g := 255 - n*contrast/2
	b := 20 + n*contrast
	return color.RGBA{R: r, G: g, B: b, A: 255}
}

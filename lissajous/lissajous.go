// Lissajous generates gif animations of random lissajous figure
package lissajous

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

const (
	paletteSize   = 100
	colorPixelRun = 120
)

func createPalette() []color.Color {
	var palette = []color.Color{color.Black}
	for i := 1; i < paletteSize; i++ {
		color := color.RGBA{
			R: uint8(rand.Intn(255)),
			G: uint8(rand.Intn(255)),
			B: uint8(rand.Intn(255)),
			A: 255}
		palette = append(palette, color)
	}
	return palette
}

func Create(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas size [-size ... + size]
		nframes = 64    // animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	palette := createPalette()
	getColorIndex := colorIndexIterator()

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			xPlot := size + int(x*size+0.5)
			yPlot := size + int(y*size+0.5)
			img.SetColorIndex(xPlot, yPlot, getColorIndex())
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)

	if err != nil {
		fmt.Fprint(out, err)
	}
}

func colorIndexIterator() func() uint8 {
	currentPixelCount := 0
	currentColorIndex := uint8(rand.Intn(paletteSize - 1))

	return func() uint8 {
		if currentPixelCount == colorPixelRun {
			currentPixelCount = 0

			currentColorIndex = uint8(rand.Intn(paletteSize - 1))
		}
		currentPixelCount++
		return currentColorIndex
	}
}

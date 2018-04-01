// Lissajous generates gif animations of random lissajous figure
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas size [-size ... + size]
		nframes = 64    // animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	//freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		// for t := 0.0; t < cycles*2*math.Pi; t += res {
		// 	x := math.Sin(t)
		// 	y := math.Sin(t*freq + phase)
		// 	xPlot := size + int(x*size+0.5)
		// 	yPlot := size + int(y*size+0.5)
		// 	//fmt.Printf("x:%v y:%v\n", xPlot, yPlot)
		// 	img.SetColorIndex(xPlot, yPlot, blackIndex)
		// }
		for i := 0; i < size; i++ {
			img.SetColorIndex(100, i, blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
}

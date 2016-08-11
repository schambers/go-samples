// Lissahous generates GIF animations of random lissajous figures
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var alpha uint8 = 0xFF
var blue uint8 = 0xBB
var green uint8 = 0xFF
var red uint8 = 0xDD

var palette = []color.Color{color.Black, color.RGBA{R: red, G: green, B: blue, A: alpha}}

const (
	blackIndex = 0
	greenIndex = 1
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nFrames = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nFrames}
	phase := 0.0 // phase difference
	for i := 0; i < nFrames; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

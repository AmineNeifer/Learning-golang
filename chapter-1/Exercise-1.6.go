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

var palette = []color.Color{color.RGBA{0xf9,0x85,0x8b,0xff}, color.RGBA{0xed, 0x33, 0x5f, 0xff}, color.RGBA{0x76, 0x11, 0x37, 0xff}}

func main() {
	lissajous(os.Stdout)
}
func lissajous(out io.Writer) {
	const (
		cycles = 5
		// number of complete x oscillator revolutions
		res  = 0.001 // angular resolution
		size = 100
		// image canvas covers [-size..+size]
		nframes = 64
		// number of animation frames
		delay = 8
	// delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			curve_color := uint8(math.Round(t))%uint8(len(palette))
			if curve_color == 0 { curve_color++}
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				curve_color)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

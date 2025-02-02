package main

import (
	"image"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		lissajous(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8003", nil))
}

func lissajous2(out io.Writer, params []string) {
	const (
		cycles  = 5		// number of complete x oscillator revolutions
		res     = 0.001	// angular resolution
		size    = 100	// image canvas covers [-size..+size]
		nframes = 64	// number of animation frames
		delay	= 8		// delay between frames in 10ms unit
	)

	frep := rand.Float64() * 3.0	// relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0	// phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles * 2 * math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t * frep + phase)
			img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5),
				uint8(rand.Uint32() % uint32(len(palette))))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)	// NOTE : ignoring encoding errors
}

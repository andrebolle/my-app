package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func ifNilPanic(err error) {
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
}

type col struct {
	r, g, b byte
}

// Window width and height
const w, h int = 800, 600

func pixel(x, y int, c col, pixels []byte) {
	i := (y*w + x) * 4

	if i < len(pixels)-4 && i >= 0 {
		pixels[i] = c.r
		pixels[i+1] = c.g
		pixels[i+2] = c.b
	}

}

func main() {

	var anywhere int32 = sdl.WINDOWPOS_UNDEFINED
	var useGPU uint32 = sdl.RENDERER_ACCELERATED
	var rgba uint32 = sdl.PIXELFORMAT_ABGR8888

	// Create a window
	win, err := sdl.CreateWindow("SDL2", anywhere, anywhere, int32(w), int32(h), sdl.WINDOW_SHOWN)
	ifNilPanic(err)
	defer win.Destroy()

	// Create a renderer
	ren, err := sdl.CreateRenderer(win, -1, useGPU)
	ifNilPanic(err)
	defer ren.Destroy()

	// Create a texture
	tex, err := ren.CreateTexture(rgba, sdl.TEXTUREACCESS_STREAMING, int32(w), int32(h))
	ifNilPanic(err)
	defer tex.Destroy()

	pixels := make([]byte, w*h*4)

	for {
		// rand.Read(pixels)
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				pixel(x, y, col{byte(x % 255), 0, 0}, pixels)
			}
		}

		tex.Update(nil, pixels, w*4)
		ren.Copy(tex, nil, nil)
		ren.Present()
	}

	sdl.Delay(2000)

}

package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

var winTitle string = "No Good"
var winWidth, winHeight int = 800, 600

func main() {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var event sdl.Event
	var running bool

	sdl.Init(sdl.INIT_VIDEO)

	window = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	if window == nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", sdl.GetError())
		os.Exit(1)
	}

	renderer = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if renderer == nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", sdl.GetError())
		os.Exit(2)
	}

	ttf.Init()
	f, err := ttf.OpenFont("assets/roboto/Roboto-Regular.ttf", 12)
	if err != nil {
		fmt.Println("font error:", err)
	}
	textsurf := f.RenderText_Solid("No Good", sdl.Color{200,200,200,128})
	fmt.Println(textsurf)

	texttex := renderer.CreateTextureFromSurface(textsurf)
	//textrect := sdl.Rect{0, 0, 128, 32}

	running = true
	for running {
		renderer.Clear()
		renderer.Copy(texttex, nil, nil)
		renderer.Present()
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
//			case *sdl.MouseMotionEvent:
//				fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
//					t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)
			case *sdl.MouseButtonEvent:
				fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
			case *sdl.MouseWheelEvent:
				fmt.Printf("[%d ms] MouseWheel\ttype:%d\tid:%d\tx:%d\ty:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y)
			case *sdl.KeyUpEvent:
				fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
					t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
			}
		}
	}

	renderer.Destroy()
	window.Destroy()
}

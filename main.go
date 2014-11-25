package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"

	"github.com/willemvds/NoGood/ui"
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
	f, err := ttf.OpenFont("assets/roboto/Roboto-Regular.ttf", 14)
	if err != nil {
		fmt.Println("font error:", err)
	}
	textsurf := f.RenderText_Blended("No Good", sdl.Color{230, 230, 230, 255})
	texttex := renderer.CreateTextureFromSurface(textsurf)
	textrect := sdl.Rect{100, 100, 56, 20}

	//textsurf2 := f.RenderText_Blended("A long long time ago. Text wasn't this ugly.", sdl.Color{0, 0, 0, 255})
	//texttex2 := renderer.CreateTextureFromSurface(textsurf2)
	//textrect2 := sdl.Rect{200, 350, 400, 20}

	uielements := make([]ui.Element, 0)
	tv := ui.NewTextView(200, 200, 100, 20)
	tv.Font = f
	uielements = append(uielements, tv)

	tv2 := ui.NewTextView(200, 300, 500, 200)
	tv2.Font = f
	uielements = append(uielements, tv2)

	button := ui.NewButton(10, 10, 80, 18)
	uielements = append(uielements, button)

	textbutton := ui.NewTextButton(10, 100, 80, 18)
	uielements = append(uielements, textbutton)

	running = true
	for running {
		renderer.Clear()
		renderer.Copy(texttex, nil, &textrect)


		for _, element := range uielements {
			element.Draw(renderer)
		}
		//renderer.Copy(texttex2, nil, &textrect2)

		renderer.Present()
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			//	case *sdl.MouseMotionEvent:
			//	fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
			//	t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)
			case *sdl.MouseButtonEvent:
				fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
				for _, element := range uielements {
					if t.X >= int32(element.X()) && t.X <= int32(element.X()+element.Width()) &&
						t.Y >= int32(element.Y()) && t.Y <= int32(element.Y()+element.Height()) {
						element.HandleEvent(event)
					}

				}
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

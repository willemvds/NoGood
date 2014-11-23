package ui

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

type TextView struct {
	Region
	Text string
}

func NewTextView(x, y, w, h int) *TextView {
	tv := TextView{
		Region: Region{x, y, w, h},
	}
	tv.Text = "A long time ago, text wasn't this ugly."
	return &tv
}

func (tv *TextView) HandleEvent(ev sdl.Event) {
	print("[textview] got event, did nothing, good day.\n")
}

func (tv *TextView) Draw(r *sdl.Renderer) {
	ttf.Init()
	rect := sdl.Rect{int32(tv.x), int32(tv.y), int32(tv.width), int32(tv.height)}
	r.SetDrawColor(0, 255, 0, 255)
	r.FillRect(&rect)
	r.SetDrawColor(0, 0, 0, 255)
	f, err := ttf.OpenFont("assets/roboto/Roboto-Regular.ttf", 14)
	if err != nil {
		return
	}
	textsurf := f.RenderText_Blended(tv.Text, sdl.Color{0, 0, 0, 255})
	texttex := r.CreateTextureFromSurface(textsurf)
	textrect := sdl.Rect{200, 350, 400, 20}
	r.Copy(texttex, nil, &textrect)
}

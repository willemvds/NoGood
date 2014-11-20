package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type TextView struct {
	Region
}

func NewTextView(x, y, w, h int) *TextView {
	tv := TextView{
		Region{x, y, w, h},
	}
	return &tv
}

func (tv *TextView) HandleEvent(ev sdl.Event) {
	print("[textview] got event, did nothing, good day.\n")
}

func (tv *TextView) Draw(r *sdl.Renderer) {
	rect := sdl.Rect{int32(tv.x), int32(tv.y), int32(tv.width), int32(tv.height)}
	r.SetDrawColor(0, 255, 0, 255)
	r.FillRect(&rect)
	r.SetDrawColor(0, 0, 0, 255)
}

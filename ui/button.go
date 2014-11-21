package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Button struct {
	Region
}

func NewButton(x, y, w, h int) *Button {
	b := Button{
		Region{x, y, w, h},
	}
	return &b
}

func (b *Button) HandleEvent(ev sdl.Event) {
	print("[button] got event, should probably do something useful now")
}

func (b *Button) Draw(r *sdl.Renderer) {
	rect := sdl.Rect{int32(b.x), int32(b.y), int32(b.width), int32(b.height)}
	r.SetDrawColor(255, 0, 0, 255)
	r.FillRect(&rect)
	r.SetDrawColor(0, 0, 0, 255)
}


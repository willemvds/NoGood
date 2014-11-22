package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type TextButton struct {
	Button
	text string
}

func NewTextButton(x, y, w, h int) *TextButton {
	tb := TextButton{
		Button: Button{
			Region{x, y, w, h},
		},
	}
	return &tb
}

func (tb *TextButton) HandleEvent(ev sdl.Event) {
	print("[textbutton] i don't care about your event dude!")
}

func (tb *TextButton) Draw(r *sdl.Renderer) {
	rect := sdl.Rect{int32(tb.x), int32(tb.y), int32(tb.width), int32(tb.height)}
	r.SetDrawColor(0, 0, 200, 255)
	r.FillRect(&rect)
	r.SetDrawColor(0, 0, 0, 255)
}

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
	print("got event, did nothing, good day. ")
}

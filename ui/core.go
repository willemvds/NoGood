package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Element interface {
	X() int
	Y() int
	Width() int
	Height() int
	//Parent() Element
	HandleEvent(sdl.Event)
}

type Region struct {
	x      int
	y      int
	width  int
	height int
}

func (r *Region) X() int {
	return r.x
}

func (r *Region) Y() int {
	return r.y
}

func (r *Region) Width() int {
	return r.width
}

func (r *Region) Height() int {
	return r.height
}

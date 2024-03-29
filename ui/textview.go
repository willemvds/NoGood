package ui

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type TextView struct {
	Region
	Text string
	Font *ttf.Font
}

func NewTextView(x, y, w, h int) *TextView {
	tv := TextView{
		Region: Region{x, y, w, h},
	}
	tv.Text = "A long time ago, text wasn't this ugly."
	return &tv
}

func (tv *TextView) HandleEvent(ev sdl.Event) {
	print("[textview] got event\n")
	switch t := ev.(type) {
	case *sdl.KeyboardEvent:
		tv.Text = tv.Text + string(t.Keysym.Sym)
	}
}

func (tv *TextView) Draw(r *sdl.Renderer) {
	rect := sdl.Rect{int32(tv.x), int32(tv.y), int32(tv.width), int32(tv.height)}
	r.SetDrawColor(0, 255, 0, 255)
	r.FillRect(&rect)
	r.SetDrawColor(0, 0, 0, 255)
	if tv.Font != nil {
		textsurf, _ := tv.Font.RenderUTF8Blended(tv.Text, sdl.Color{0, 0, 0, 255})
		texttex, _ := r.CreateTextureFromSurface(textsurf)
		textrect := sdl.Rect{200, 350, 400, 20}
		r.Copy(texttex, nil, &textrect)
	}
}

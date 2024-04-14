package ecs

import (
	"desktop-buddy/internal/core"
	"desktop-buddy/pkg/helpers"

	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	alpha      float32
	frameIndex int

	FrameLength, FrameTime      int
	Height, Width, X, Y, Vx, Vy float64
	GifName                     string
	SingleAsset                 bool
	Invert                      bool
	AnchorMargin                float64
}

func (e *Entity) Update() {
	e.Y = float64(core.ScreenHeight) - e.Height
	e.Y += e.Vy

	if core.FrameTick%e.FrameTime == 0 {
		e.frameIndex = (e.frameIndex + 1) % e.FrameLength
	}

	if e.alpha < 1 {
		e.alpha += 0.01
	}
}

func (e *Entity) Draw(screen *ebiten.Image, gif *helpers.CustomGif) {
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(e.alpha)
	if e.SingleAsset && e.Invert {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(e.Width), 0)
	}
	op.GeoM.Translate(e.X, e.Y-e.AnchorMargin)
	screen.DrawImage(gif.Frames[e.frameIndex], op)
}

func (e *Entity) SetGif(name string, gif *helpers.CustomGif) {
	e.GifName = name
	e.frameIndex = 0
	e.Height = gif.Height
	e.Width = gif.Width
	e.FrameLength = gif.Length
}

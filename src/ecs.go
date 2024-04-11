package main

import (
	"desktop-buddy/src/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	x, y, vx, vy          float64
	frameIndex, frameTime int
	gifName               string
	singleAsset           bool
	alpha                 float32
	invert                bool
	// anchor                string
	anchorMargin float64
}

func (e *Entity) update() {
	e.y = float64(screenHeight) - assets.LoadedGifs[e.gifName].Height
	e.y += e.vy

	if frameCount%e.frameTime == 0 {
		e.frameIndex = (e.frameIndex + 1) % assets.LoadedGifs[e.gifName].Length
	}

	if e.alpha < 1 {
		e.alpha += 0.01
	}
}

func (e *Entity) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(e.alpha)
	if e.singleAsset && e.invert {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(assets.LoadedGifs[e.gifName].Width), 0)
	}
	op.GeoM.Translate(e.x, e.y-e.anchorMargin)
	screen.DrawImage(assets.LoadedGifs[e.gifName].Frames[e.frameIndex], op)
}

func (e *Entity) setGif(name string) {
	e.gifName = name
	e.frameIndex = 0
}

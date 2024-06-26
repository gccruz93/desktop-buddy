/*
Copyright © 2024 <https://github.com/gccruz93> twpax
*/
//go:generate goversioninfo
package main

import (
	"bytes"
	"desktop-buddy/assets"
	"desktop-buddy/internal/core"
	"desktop-buddy/internal/emotes"
	"desktop-buddy/internal/mobs"
	"desktop-buddy/pkg/helpers"
	_ "embed"
	"image"
	_ "image/png"
	"log"

	"github.com/energye/systray"
	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	core.Cfg.Load()
	mobs.LoadConfig()
	emotes.LoadConfig()
}

type Game struct{}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Update() error {
	core.FrameTick++

	mobs.Update()
	emotes.Update()

	if core.FrameTick%emotes.NextSpawnTick == 0 && emotes.EmoteActive == nil {
		emotes.SpawnRandom()
		mobs.List[helpers.Random(0, len(mobs.List)-1)].Emote = emotes.EmoteActive
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	mobs.Draw(screen)
	emotes.Draw(screen)
}

func setScreenArea() {
	sw, sh := ebiten.ScreenSizeInFullscreen()
	screenHeight := sh - core.Cfg.ScreenPaddingBottom
	screenWidth := sw * core.Cfg.ScreenMonitors
	ebiten.SetWindowSize(screenWidth, screenHeight)
	core.ScreenHeight = float64(screenHeight)
	core.ScreenWidth = float64(screenWidth)
}

func main() {
	ebiten.SetWindowTitle(core.Title)
	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowFloating(true)
	ebiten.SetWindowMousePassthrough(true)
	setScreenArea()

	img, _, err := image.Decode(bytes.NewReader(assets.Icon))
	if err != nil {
		log.Fatal(err)
	}
	iconImages := []image.Image{img}
	ebiten.SetWindowIcon(iconImages)

	op := &ebiten.RunGameOptions{}
	op.ScreenTransparent = true
	op.SkipTaskbar = core.Cfg.SkipTaskbar

	trayStart, trayEnd := systray.RunWithExternalLoop(onReady, onExit)
	trayStart()
	if err := ebiten.RunGameWithOptions(&Game{}, op); err != nil {
		trayEnd()
		log.Fatal(err)
	}
}

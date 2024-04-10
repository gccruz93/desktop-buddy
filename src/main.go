/*
Copyright © 2024 <https://github.com/gccruz93> twpax
*/
//go:generate goversioninfo
package main

import (
	"bytes"
	"desktop-buddy/assets"
	_ "embed"
	"image"
	_ "image/png"
	"log"

	"github.com/energye/systray"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	screenHeight = 0
	screenWidth  = 0
	title        = "Desktop Buddy"
	frameCount   = 0
	nextSpawn    = 1
)

func init() {
	cfg.Load()
	loadMobsConfig()
	assets.LoadedGifs = make(map[string]*assets.CustomGif)
}

type Game struct{}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Update() error {
	frameCount++

	if frameCount%(nextSpawn*ebiten.TPS()) == 0 && len(mobs) < cfg.MobsSpawnTotal {
		SpawnRandom(1)
	}

	mobsAlive := mobs[:0]
	for _, e := range mobs {
		e.Update()
		if cfg.MobsSpawnCycle {
			e.lifeTime--
		}

		if e.lifeTime > 0 {
			mobsAlive = append(mobsAlive, e)
		}
	}
	mobs = mobsAlive

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, e := range mobs {
		e.Draw(screen)
	}
}

func setScreenArea() {
	sw, sh := ebiten.ScreenSizeInFullscreen()
	screenHeight = sh - cfg.ScreenPaddingBottom
	screenWidth = sw * cfg.ScreenMonitors
	ebiten.SetWindowSize(screenWidth, screenHeight)
}

func main() {
	ebiten.SetWindowTitle(title)
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
	op.SkipTaskbar = cfg.SkipTaskbar

	trayStart, trayEnd := systray.RunWithExternalLoop(onReady, onExit)
	trayStart()
	if err := ebiten.RunGameWithOptions(&Game{}, op); err != nil {
		trayEnd()
		log.Fatal(err)
	}
}

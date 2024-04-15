package mobs

import (
	"desktop-buddy/internal/core"
	"desktop-buddy/internal/emotes"
	"desktop-buddy/pkg/helpers"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	List       []*Mob
	rarityList []int
)

var (
	cachedConfig []*MobConfig
	cachedGifs   map[string]*helpers.CustomGif
)

var nextSpawnTick = 1

type Mob struct {
	Emote                         *emotes.Emote
	assetName                     string
	speed                         float64
	moveFuel, idleTime, idleCount int
	idleFrametime, walkFrametime  int
	status                        string // "idle", "walk"
	Immortal                      bool
	LifeTime                      int

	alpha                       float32
	frameIndex                  int
	FrameLength, FrameTime      int
	Height, Width, X, Y, Vx, Vy float64
	GifName                     string
	SingleAsset                 bool
	Invert                      bool
	AnchorMargin                float64
}

func (m *Mob) Update() {
	if m.speed > 0 {
		if m.moveFuel > 0 {
			m.X += m.Vx
			if m.X+float64(m.Width) >= float64(core.ScreenWidth) {
				m.Vx = -m.speed
				m.Invert = false
			} else if m.X <= 1 {
				m.Vx = m.speed
				m.Invert = true
			}
			m.moveFuel--
		} else {
			if m.idleCount == 0 {
				m.setIdle()
				loops := helpers.Random(8, 12)
				m.idleTime = m.FrameTime * m.FrameLength * loops
			}

			m.idleCount++

			if m.idleCount >= m.idleTime {
				m.idleCount = 0
				m.idleTime = 0

				if helpers.Random(0, 1) == 0 {
					m.setWalkLeft()
				} else {
					m.setWalkRight()
				}

				steps := helpers.Random(1, 6)
				m.moveFuel = m.FrameTime*m.FrameLength*steps - m.FrameTime
			}
		}
	}

	m.Y = float64(core.ScreenHeight) - m.Height
	m.Y += m.Vy

	if core.FrameTick%m.FrameTime == 0 {
		m.frameIndex = (m.frameIndex + 1) % m.FrameLength
	}

	if m.alpha < 1 {
		m.alpha += 0.01
	}
}
func (m *Mob) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(m.alpha)
	if m.SingleAsset && m.Invert {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(m.Width), 0)
	}
	op.GeoM.Translate(m.X, m.Y-m.AnchorMargin)
	screen.DrawImage(cachedGifs[m.GifName].Frames[m.frameIndex], op)
}

func (m *Mob) setStatus(status string) {
	m.status = status
	if m.SingleAsset {
		gifName := m.assetName + "_" + m.status
		m.setGif(gifName)
	} else {
		posfix := "_left"
		if m.Invert {
			posfix = "_right"
		}
		gifName := m.assetName + "_" + m.status + posfix
		m.setGif(gifName)
	}
}
func (m *Mob) setIdle() {
	m.FrameTime = m.idleFrametime
	m.Vx = 0
	m.setStatus("idle")
}
func (m *Mob) setWalk() {
	m.FrameTime = m.walkFrametime
	m.setStatus("walk")
}
func (m *Mob) setWalkLeft() {
	m.Vx = -m.speed
	m.Invert = false
	m.setWalk()
}
func (m *Mob) setWalkRight() {
	m.Vx = m.speed
	m.Invert = true
	m.setWalk()
}
func (m *Mob) setSpawn() {
	m.X = float64(helpers.Random(0, int(float64(core.Cfg.ScreenMonitors)*core.ScreenWidth-2*m.Width)))
	m.LifeTime = helpers.Random(core.Cfg.MobsDespawnSecondsMin, core.Cfg.MobsDespawnSecondsMin) * ebiten.TPS()
	m.setIdle()
	if helpers.Random(0, 1) == 0 {
		m.Invert = false
	} else {
		m.Invert = true
	}
}
func (m *Mob) setGif(name string) {
	m.GifName = name
	m.frameIndex = 0
	m.Height = cachedGifs[name].Height
	m.Width = cachedGifs[name].Width
	m.FrameLength = cachedGifs[name].Length
}
func (m *Mob) loadGif(status string) {
	a := loadAsset(m.assetName + "_" + status)
	if a {
		m.SingleAsset = true
	} else {
		loadAsset(m.assetName + "_" + status + "_left")
		loadAsset(m.assetName + "_" + status + "_right")
	}
}

func (m *Mob) setupIdle(frameTime int) {
	m.loadGif("idle")
	m.idleFrametime = frameTime
}
func (m *Mob) setupWalk(frameTime int) {
	m.loadGif("walk")
	m.walkFrametime = frameTime
}

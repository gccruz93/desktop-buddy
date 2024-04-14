package mobs

import (
	"desktop-buddy/internal/core"
	"desktop-buddy/internal/ecs"
	"desktop-buddy/pkg/helpers"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	MobsAlive  []*Mob
	MobsRarity []int
)

var (
	CachedMobs     []*MobsConfig
	CachedMobsGifs map[string]*helpers.CustomGif
)

type Mob struct {
	ecs.Entity
	assetName                     string
	speed                         float64
	moveFuel, idleTime, idleCount int
	idleFrametime, walkFrametime  int
	status                        string // "idle", "walk"

	Immortal bool
	LifeTime int
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

	m.Entity.Update()
}
func (m *Mob) Draw(screen *ebiten.Image) {
	m.Entity.Draw(screen, CachedMobsGifs[m.GifName])
}

package mobs

import (
	"desktop-buddy/internal/core"
	"desktop-buddy/pkg/helpers"

	"github.com/hajimehoshi/ebiten/v2"
)

func SpawnRandom(n int) int {
	nextSpawn := helpers.Random(core.Cfg.MobsSpawnSecondsMin, core.Cfg.MobsSpawnSecondsMax)

	if len(CachedMobs) == 0 {
		return nextSpawn
	}

	for n > 0 {
		mobConfig := CachedMobs[MobsRarity[helpers.Random(0, len(MobsRarity)-1)]]
		mob := &Mob{
			speed:     mobConfig.Speed,
			Immortal:  mobConfig.Immortal,
			assetName: mobConfig.Name,
		}
		mob.AnchorMargin = mobConfig.AnchorMargin

		mob.setupIdle(mobConfig.IdleFrametime)
		mob.setupWalk(mobConfig.WalkFrametime)

		mob.setStatus("idle")
		mob.setSpawn()

		MobsAlive = append(MobsAlive, mob)
		n--
	}

	return nextSpawn
}

func (m *Mob) setStatus(status string) {
	m.status = status
	if m.SingleAsset {
		gifName := m.assetName + "_" + m.status
		m.SetGif(gifName, CachedMobsGifs[gifName])
	} else {
		posfix := "_left"
		if m.Invert {
			posfix = "_right"
		}
		gifName := m.assetName + "_" + m.status + posfix
		m.SetGif(gifName, CachedMobsGifs[gifName])
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
func (m *Mob) loadGif(status string) {
	a := loadMob(m.assetName + "_" + status)
	if a {
		m.SingleAsset = true
	} else {
		loadMob(m.assetName + "_" + status + "_left")
		loadMob(m.assetName + "_" + status + "_right")
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

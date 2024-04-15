package mobs

import (
	"desktop-buddy/internal/core"

	"github.com/hajimehoshi/ebiten/v2"
)

func Update() {
	eventSpawnRandom()

	mobsAlive := List[:0]
	for _, e := range List {
		e.Update()
		if core.Cfg.MobsSpawnCycle && !e.Immortal {
			e.LifeTime--
		}

		if e.Emote != nil {
			e.Emote.X = e.X
			e.Emote.Y = e.Y
		}

		if e.LifeTime > 0 {
			mobsAlive = append(mobsAlive, e)
		}
	}
	List = mobsAlive
}

func Draw(screen *ebiten.Image) {
	for _, e := range List {
		e.Draw(screen)
	}
}

func eventSpawnRandom() {
	if core.FrameTick%nextSpawnTick == 0 && len(List) < core.Cfg.MobsSpawnTotal {
		SpawnRandom(1)
	}
}

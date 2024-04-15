package mobs

import (
	"bytes"
	"desktop-buddy/internal/core"
	"desktop-buddy/pkg/helpers"
	"image/gif"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func loadAsset(name string) bool {
	if _, ok := cachedGifs[name]; ok {
		return true
	}
	file, err := os.ReadFile("./assets/mobs/" + name + ".gif")
	if err != nil {
		return false
	}
	loadedGif, _ := gif.DecodeAll(bytes.NewReader(file))
	cachedGifs[name] = helpers.SplitAnimatedGIF(loadedGif)
	return true
}

func SpawnRandom(n int) {
	nextSpawnTick = helpers.Random(core.Cfg.MobsSpawnSecondsMin, core.Cfg.MobsSpawnSecondsMax) * ebiten.TPS()

	if len(cachedConfig) == 0 {
		return
	}

	for n > 0 {
		mobConfig := cachedConfig[rarityList[helpers.Random(0, len(rarityList)-1)]]
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

		List = append(List, mob)
		n--
	}
}

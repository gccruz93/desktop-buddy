package emotes

import (
	"bytes"
	"desktop-buddy/internal/core"
	"desktop-buddy/pkg/helpers"
	"image/gif"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func loadAsset(name string) {
	file, err := os.ReadFile("./assets/emotes/" + name + ".gif")
	if err != nil {
		return
	}
	loadedGif, _ := gif.DecodeAll(bytes.NewReader(file))
	EmoteActive.gif = helpers.SplitAnimatedGIF(loadedGif)
}

func SpawnRandom() {
	NextSpawnTick = helpers.Random(core.Cfg.EmoteSpawnSecondsMin, core.Cfg.EmoteSpawnSecondsMax) * ebiten.TPS()

	if len(cachedConfig) == 0 {
		return
	}

	emoteConfig := cachedConfig[rarityList[helpers.Random(0, len(rarityList)-1)]]
	EmoteActive = &Emote{
		anchorMargin: emoteConfig.AnchorMargin,
		frameTime:    emoteConfig.Frametime,
		loops:        emoteConfig.Loops,
	}
	loadAsset(emoteConfig.Name)
}

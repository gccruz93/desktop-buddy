package mobs

import (
	"bytes"
	"desktop-buddy/pkg/helpers"
	"encoding/json"
	"image/gif"
	"log"
	"os"
)

type MobsConfig struct {
	Name          string  `json:"name"`
	Speed         float64 `json:"speed"`
	IdleFrametime int     `json:"idle_frametime"`
	WalkFrametime int     `json:"walk_frametime"`
	Rarity        int     `json:"rarity"`
	AnchorMargin  float64 `json:"anchor_margin"`
	Immortal      bool    `json:"immortal"`
}

func loadMob(name string) bool {
	if _, ok := CachedMobsGifs[name]; ok {
		return true
	}
	file, err := os.ReadFile("./assets/mobs/" + name + ".gif")
	if err != nil {
		return false
	}
	loadedGif, _ := gif.DecodeAll(bytes.NewReader(file))
	CachedMobsGifs[name] = helpers.SplitAnimatedGIF(loadedGif)
	return true
}

func LoadMobsConfig() {
	MobsAlive = nil
	MobsRarity = nil
	CachedMobs = nil
	CachedMobsGifs = nil
	CachedMobsGifs = make(map[string]*helpers.CustomGif)

	var mobs []*MobsConfig

	file, err := os.Open("mobs.json")
	if err != nil {
		if os.IsNotExist(err) {
			mob := &MobsConfig{
				Name:          "poring",
				Speed:         1.3,
				IdleFrametime: 6,
				WalkFrametime: 3,
				Rarity:        1,
				AnchorMargin:  1,
				Immortal:      false,
			}
			mobs = append(mobs, mob)
			createMobsConfig(mobs)
		} else {
			panic(err)
		}
	} else {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&mobs)
		if err != nil {
			panic(err)
		}
	}

	for i, mob := range mobs {
		CachedMobs = append(CachedMobs, mob)
		for j := 0; j <= mob.Rarity; j++ {
			MobsRarity = append(MobsRarity, i)
		}
	}

	defer file.Close()
}
func createMobsConfig(mobs []*MobsConfig) {
	file, err := os.Create("mobs.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(mobs)
	if err != nil {
		log.Fatal(err)
	}
}

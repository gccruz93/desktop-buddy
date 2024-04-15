package mobs

import (
	"desktop-buddy/pkg/helpers"
	"encoding/json"
	"log"
	"os"
)

type MobConfig struct {
	Name          string  `json:"name"`
	Speed         float64 `json:"speed"`
	IdleFrametime int     `json:"idle_frametime"`
	WalkFrametime int     `json:"walk_frametime"`
	Rarity        int     `json:"rarity"`
	AnchorMargin  float64 `json:"anchor_margin"`
	Immortal      bool    `json:"immortal"`
}

func LoadConfig() {
	List = nil
	rarityList = nil
	cachedConfig = nil
	cachedGifs = nil
	cachedGifs = make(map[string]*helpers.CustomGif)
	nextSpawnTick = 1

	var mobs []*MobConfig

	file, err := os.Open("mobs.json")
	if err != nil {
		if os.IsNotExist(err) {
			mob := &MobConfig{
				Name:          "poring",
				Speed:         1.3,
				IdleFrametime: 6,
				WalkFrametime: 3,
				Rarity:        1,
				AnchorMargin:  1,
				Immortal:      false,
			}
			mobs = append(mobs, mob)
			createConfig(mobs)
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
		cachedConfig = append(cachedConfig, mob)
		for j := 0; j <= mob.Rarity; j++ {
			rarityList = append(rarityList, i)
		}
	}

	defer file.Close()
}

func createConfig(data []*MobConfig) {
	file, err := os.Create("mobs.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}

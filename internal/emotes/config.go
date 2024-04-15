package emotes

import (
	"encoding/json"
	"log"
	"os"
)

type EmoteConfig struct {
	Name         string  `json:"name"`
	Frametime    int     `json:"frametime"`
	Rarity       int     `json:"rarity"`
	AnchorMargin float64 `json:"anchor_margin"`
	Loops        int     `json:"loops"`
}

func LoadConfig() {
	rarityList = nil
	cachedConfig = nil

	var emotes []*EmoteConfig

	file, err := os.Open("emotes.json")
	if err != nil {
		if os.IsNotExist(err) {
			mob := &EmoteConfig{
				Name:         "flag_brazil",
				Frametime:    6,
				Rarity:       1,
				AnchorMargin: 10,
				Loops:        0,
			}
			emotes = append(emotes, mob)

			mob = &EmoteConfig{
				Name:         "heh",
				Frametime:    3,
				Rarity:       1,
				AnchorMargin: 10,
				Loops:        1,
			}
			emotes = append(emotes, mob)

			mob = &EmoteConfig{
				Name:         "hunf2",
				Frametime:    4,
				Rarity:       1,
				AnchorMargin: 10,
				Loops:        1,
			}
			emotes = append(emotes, mob)

			mob = &EmoteConfig{
				Name:         "idea",
				Frametime:    4,
				Rarity:       1,
				AnchorMargin: 10,
				Loops:        1,
			}
			emotes = append(emotes, mob)

			mob = &EmoteConfig{
				Name:         "interrogation",
				Frametime:    4,
				Rarity:       1,
				AnchorMargin: 10,
				Loops:        1,
			}
			emotes = append(emotes, mob)

			mob = &EmoteConfig{
				Name:         "no",
				Frametime:    4,
				Rarity:       1,
				AnchorMargin: 10,
				Loops:        1,
			}
			emotes = append(emotes, mob)

			mob = &EmoteConfig{
				Name:         "sing",
				Frametime:    4,
				Rarity:       1,
				AnchorMargin: 10,
				Loops:        1,
			}
			emotes = append(emotes, mob)

			createConfig(emotes)
		} else {
			panic(err)
		}
	} else {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&emotes)
		if err != nil {
			panic(err)
		}
	}

	for i, mob := range emotes {
		cachedConfig = append(cachedConfig, mob)
		for j := 0; j <= mob.Rarity; j++ {
			rarityList = append(rarityList, i)
		}
	}

	defer file.Close()
}

func createConfig(data []*EmoteConfig) {
	file, err := os.Create("emotes.json")
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

package main

import (
	"desktop-buddy/assets"
	"encoding/json"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Mob struct {
	Entity
	assetName                               string
	speed                                   float64
	moveFuel, idleTime, idleCount, lifeTime int

	idleFrametime int
	walkFrametime int

	status string // "idle", "walk"
}

var mobs []*Mob

func (m *Mob) Update() {
	if m.speed > 0 {
		if m.moveFuel > 0 {
			m.x += m.vx
			if m.x+float64(assets.LoadedGifs[m.gifName].Width) >= float64(screenWidth) {
				m.vx = -m.speed
				m.invert = false
			} else if m.x <= 1 {
				m.vx = m.speed
				m.invert = true
			}
			m.moveFuel--
		} else {
			if m.idleCount == 0 {
				m.setIdle()
				loops := random(8, 12)
				m.idleTime = m.frameTime * assets.LoadedGifs[m.gifName].Length * loops
			}

			m.idleCount++

			if m.idleCount >= m.idleTime {
				m.idleCount = 0
				m.idleTime = 0

				if random(0, 1) == 0 {
					m.setWalkLeft()
				} else {
					m.setWalkRight()
				}

				steps := random(1, 6)
				m.moveFuel = m.frameTime*assets.LoadedGifs[m.gifName].Length*steps - m.frameTime
			}
		}
	}

	m.Entity.update()
}
func (m *Mob) Draw(screen *ebiten.Image) {
	m.Entity.draw(screen)
}
func (m *Mob) setStatus(status string) {
	m.status = status
	if m.singleAsset {
		m.setGif(m.assetName + "_" + m.status)
	} else {
		posfix := "_left"
		if m.invert {
			posfix = "_right"
		}
		m.setGif(m.assetName + "_" + m.status + posfix)
	}
}
func (m *Mob) setIdle() {
	m.frameTime = m.idleFrametime
	m.vx = 0
	m.setStatus("idle")
}
func (m *Mob) setWalk() {
	m.frameTime = m.walkFrametime
	m.setStatus("walk")
}
func (m *Mob) setWalkLeft() {
	m.vx = -m.speed
	m.invert = false
	m.setWalk()
}
func (m *Mob) setWalkRight() {
	m.vx = m.speed
	m.invert = true
	m.setWalk()
}
func (m *Mob) setSpawn() {
	m.x = float64(random(0, cfg.ScreenMonitors*screenWidth-2*int(assets.LoadedGifs[m.gifName].Width)))
	m.lifeTime = random(cfg.MobsDespawnSecondsMin, cfg.MobsDespawnSecondsMin) * ebiten.TPS()
	m.setIdle()
	if random(0, 1) == 0 {
		m.invert = false
	} else {
		m.invert = true
	}
}
func (m *Mob) loadGif(status string) {
	a := assets.LoadGif(m.assetName + "_" + status)
	if a {
		m.singleAsset = true
	} else {
		assets.LoadGif(m.assetName + "_" + status + "_left")
		assets.LoadGif(m.assetName + "_" + status + "_right")
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

type MobConfig struct {
	Name          string  `json:"name"`
	Speed         float64 `json:"speed"`
	IdleFrametime int     `json:"idle_frametime"`
	WalkFrametime int     `json:"walk_frametime"`
	Rarity        int     `json:"rarity"`
	// Anchor        string  `json:"anchor"`
	AnchorMargin float64 `json:"anchor_margin"`
}

var mobsLoaded map[string]*MobConfig
var mobsRarity []string

func loadMobsConfig() {
	mobs = nil
	mobsLoaded = nil
	mobsLoaded = make(map[string]*MobConfig)

	var aux []*MobConfig

	file, err := os.Open("mobs.json")
	if err != nil {
		if os.IsNotExist(err) {
			mob := &MobConfig{
				Name:          "poring",
				Speed:         1.3,
				IdleFrametime: 6,
				WalkFrametime: 3,
				Rarity:        1,
				// Anchor:        "bottom",
				AnchorMargin: 1,
			}
			aux = append(aux, mob)
			createMobsConfig(aux)
		} else {
			panic(err)
		}
	} else {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&aux)
		if err != nil {
			panic(err)
		}
	}

	for _, mob := range aux {
		mobsLoaded[mob.Name] = mob
		for i := 0; i <= mob.Rarity; i++ {
			mobsRarity = append(mobsRarity, mob.Name)
		}
	}

	defer file.Close()
}
func createMobsConfig(config []*MobConfig) {
	file, err := os.Create("mobs.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(config)
	if err != nil {
		log.Fatal(err)
	}
}

func SpawnRandom(n int) {
	if len(mobsLoaded) == 0 {
		return
	}

	nextSpawn = random(cfg.MobsSpawnSecondsMin, cfg.MobsSpawnSecondsMax)

	for n > 0 {
		loadedMob := mobsLoaded[mobsRarity[random(0, len(mobsRarity)-1)]]
		newMob := &Mob{
			speed: loadedMob.Speed,
		}
		newMob.assetName = loadedMob.Name
		// newMob.anchor = loadedMob.Anchor
		// if loadedMob.Anchor == "bottom" {
		// 	newMob.anchorMargin = -loadedMob.AnchorMargin
		// } else {
		// }
		newMob.anchorMargin = loadedMob.AnchorMargin

		newMob.setupIdle(loadedMob.IdleFrametime)
		newMob.setupWalk(loadedMob.WalkFrametime)

		newMob.setStatus("idle")
		newMob.setSpawn()

		mobs = append(mobs, newMob)
		n--
	}
}

package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Mob struct {
	Entity
	speed                                   float64
	moveFuel, idleTime, idleCount, lifeTime int

	idleFrametime int
	walkFrametime int
	gifName       string
}

var mobs []*Mob

func (m *Mob) Update() {
	if m.moveFuel > 0 {
		m.x += m.vx
		if m.x+float64(m.width) >= float64(screenWidth) {
			m.SetWalkLeft()
		} else if m.x <= 1 {
			m.SetWalkRight()
		}
		m.moveFuel--
	} else {
		if m.idleCount == 0 {
			m.SetIdle()
			m.idleTime = random(200, 400)
		}

		m.idleCount++

		if m.idleCount >= m.idleTime {
			m.idleCount = 0
			m.idleTime = 0

			if random(0, 1) == 0 {
				m.SetWalkLeft()
			} else {
				m.SetWalkRight()
			}

			steps := random(1, 8)
			m.moveFuel = (m.frameSpeed * m.frameLength * steps) - steps
		}
	}

	m.Entity.Update()
}
func (m *Mob) Draw(screen *ebiten.Image) {
	m.Entity.Draw(screen)
}
func (m *Mob) SetIdle() {
	if m.vx > 0 {
		m.SetIdleRight()
	} else {
		m.SetIdleLeft()
	}
}
func (m *Mob) SetIdleLeft() {
	m.SetGif(m.gifName+"_idle_left", m.idleFrametime)
	m.vx = 0
}
func (m *Mob) SetIdleRight() {
	m.SetGif(m.gifName+"_idle_right", m.idleFrametime)
	m.vx = 0
}
func (m *Mob) SetWalkLeft() {
	m.SetGif(m.gifName+"_walk_left", m.walkFrametime)
	m.vx = -m.speed
}
func (m *Mob) SetWalkRight() {
	m.SetGif(m.gifName+"_walk_right", m.walkFrametime)
	m.vx = m.speed
}
func (m *Mob) SetSpawn() {
	m.x = float64(random(0, cfg.ScreenMonitors*screenWidth-2*m.width))
}
func (m *Mob) SetIdleFrametime(frametime int) {
	m.idleFrametime = frametime
}
func (m *Mob) SetWalkFrametime(frametime int) {
	m.walkFrametime = frametime
}

type MobConfig struct {
	Name          string  `json:"name"`
	Speed         float64 `json:"speed"`
	IdleFrametime int     `json:"idle_frametime"`
	WalkFrametime int     `json:"walk_frametime"`
	Rarity        int     `json:"rarity"`
}

var mobsConfig map[string]*MobConfig
var mobsRarity []string

func loadMobsConfig() {
	mobs = nil
	mobsConfig = nil
	mobsConfig = make(map[string]*MobConfig)

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
		mobsConfig[mob.Name] = mob
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
	if len(mobsConfig) == 0 {
		return
	}

	nextSpawn = random(cfg.MobsSpawnSecondsMin, cfg.MobsSpawnSecondsMax)

	for n > 0 {
		loadedMob := mobsConfig[mobsRarity[random(0, len(mobsRarity)-1)]]
		newMob := &Mob{
			speed:         loadedMob.Speed,
			gifName:       loadedMob.Name,
			idleFrametime: loadedMob.IdleFrametime,
			walkFrametime: loadedMob.WalkFrametime,
		}
		newMob.SetSpawn()
		newMob.lifeTime = random(cfg.MobsDespawnSecondsMin, cfg.MobsDespawnSecondsMin) * ebiten.TPS()
		mobs = append(mobs, newMob)
		n--
	}
}

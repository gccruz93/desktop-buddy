package main

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

type Cfg struct {
	SkipTaskbar           bool `ini:"skip_taskbar" json:"skip_taskbar"`
	ScreenPaddingBottom   int  `ini:"screen_padding_bottom" json:"screen_padding_bottom"`
	ScreenMonitors        int  `ini:"screen_monitors" json:"screen_monitors"`
	MobsCycle             bool `ini:"mobs_cycle" json:"mobs_cycle"`
	MobsSpawnMax          int  `ini:"mobs_spawn_max" json:"mobs_spawn_max"`
	MobsSpawnSecondsMin   int  `ini:"mobs_spawn_seconds_min" json:"mobs_spawn_seconds_min"`
	MobsSpawnSecondsMax   int  `ini:"mobs_spawn_seconds_max" json:"mobs_spawn_seconds_max"`
	MobsDespawnSecondsMin int  `ini:"mobs_despawn_seconds_min" json:"mobs_despawn_seconds_min"`
	MobsDespawnSecondsMax int  `ini:"mobs_despawn_seconds_max" json:"mobs_despawn_seconds_max"`
}

var cfg Cfg

func (c *Cfg) Load() {
	ini.PrettyFormat = false
	c.LoadDefaults()

	cfg, err := ini.Load("cfg.ini")
	if err != nil {
		cfg = ini.Empty()
		_ = cfg.ReflectFrom(c)
		_ = cfg.SaveTo("cfg.ini")
	}

	err = cfg.MapTo(&c)
	if err != nil {
		fmt.Printf("Fail to map file: %v", err)
		os.Exit(1)
	}

	c.ScreenMonitors = c.IntPositive(c.ScreenMonitors, 1)
	c.MobsSpawnMax = c.IntPositive(c.MobsSpawnMax, 6)
	c.MobsSpawnSecondsMin = c.IntPositive(c.MobsSpawnSecondsMin, 5)
	c.MobsSpawnSecondsMax = c.IntPositive(c.MobsSpawnSecondsMax, 20)

	c.Save()
}

func (c *Cfg) Save() {
	cfg := ini.Empty()
	_ = cfg.ReflectFrom(c)
	err := cfg.SaveTo("cfg.ini")
	if err != nil {
		fmt.Printf("Fail to save file: %v", err)
		os.Exit(1)
	}
}

func (c *Cfg) LoadDefaults() {
	c.ScreenPaddingBottom = 62
	c.ScreenMonitors = 1
	c.MobsCycle = true
	c.MobsSpawnMax = 6
	c.MobsSpawnSecondsMin = 3
	c.MobsSpawnSecondsMax = 10
	c.MobsDespawnSecondsMin = 20
	c.MobsDespawnSecondsMax = 40
}

func (c *Cfg) IntRange(val, min, max, dfault int) int {
	if val < min {
		return dfault
	} else if val > max {
		return dfault
	}
	return val
}
func (c *Cfg) IntPositive(val, dfault int) int {
	if val <= 0 {
		return dfault
	}
	return val
}
func (c *Cfg) FloatRange(val, min, max, dfault float64) float64 {
	if val < min {
		return dfault
	} else if val > max {
		return dfault
	}
	return val
}

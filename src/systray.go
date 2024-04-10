package main

import (
	"desktop-buddy/assets"
	"fmt"

	"github.com/energye/systray"
)

func onReady() {
	systray.SetIcon(assets.Icontray)
	systray.SetTitle(title)
	systray.SetTooltip(title)
	systray.SetOnRClick(func(menu systray.IMenu) {
		menu.ShowMenu()
	})

	systray.AddMenuItem("v0.2.0", "v0.2.0").Disable()

	systray.AddMenuItem("Close", "Close").Click(func() {
		cfg.Save()
		systray.Quit()
	})

	systray.AddSeparator()

	mTaskbar := systray.AddMenuItem("Hide from taskbar", "Hide from taskbar")
	mTaskbar.Click(func() {
		cfg.SkipTaskbar = !cfg.SkipTaskbar
		if cfg.SkipTaskbar {
			mTaskbar.Check()
		} else {
			mTaskbar.Uncheck()
		}
		cfg.Save()
	})
	if cfg.SkipTaskbar {
		mTaskbar.Check()
	}

	systray.AddMenuItem("Reload", "Reload").Click(func() {
		cfg.Load()
		loadMobsConfig()
		assets.ClearGifs()
		nextSpawn = 1
	})

	systray.AddSeparator()

	/**
	* ========== MOBS ==========
	 */
	mMobs := systray.AddMenuItem("Mobs", "Mobs")
	mMobsSpawnTotal := mMobs.AddSubMenuItem(fmt.Sprintf("Total: %d", cfg.MobsSpawnTotal), "")
	mMobsSpawnTotal.Disable()

	mMobs.AddSubMenuItem("Total++", "").Click(func() {
		cfg.MobsSpawnTotal++
		mMobsSpawnTotal.SetTitle(fmt.Sprintf("Total: %d", cfg.MobsSpawnTotal))
		cfg.Save()
	})

	mMobs.AddSubMenuItem("Total--", "").Click(func() {
		cfg.MobsSpawnTotal--
		mMobsSpawnTotal.SetTitle(fmt.Sprintf("Total: %d", cfg.MobsSpawnTotal))
		cfg.Save()
	})

	mMobsSpawnCycle := mMobs.AddSubMenuItem("Enable cycle", "Enable cycle")
	mMobsSpawnCycle.Click(func() {
		cfg.MobsSpawnCycle = !cfg.MobsSpawnCycle
		if cfg.MobsSpawnCycle {
			mMobsSpawnCycle.Check()
		} else {
			mMobsSpawnCycle.Uncheck()
		}
		cfg.Save()
	})
	if cfg.MobsSpawnCycle {
		mMobsSpawnCycle.Check()
	}

	systray.AddMenuItem("Spawn", "Spawn").Click(func() {
		SpawnRandom(1)
	})
}

func onExit() {}

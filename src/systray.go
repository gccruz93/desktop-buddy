package main

import (
	"fmt"

	"github.com/energye/systray"
	"github.com/hajimehoshi/ebiten/v2"
)

func onReady() {
	systray.SetIcon(icontray)
	systray.SetTitle(title)
	systray.SetTooltip(title)
	systray.SetOnRClick(func(menu systray.IMenu) {
		menu.ShowMenu()
	})

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

	/**
	* ========== MOBS ==========
	 */
	mMobs := systray.AddMenuItem("Mobs", "Mobs")
	mMobsMaximo := mMobs.AddSubMenuItem(fmt.Sprintf("Total: %d", cfg.MobsSpawnMax), "")
	mMobsMaximo.Disable()

	mMobs.AddSubMenuItem("Total++", "").Click(func() {
		cfg.MobsSpawnMax++
		mMobsMaximo.SetTitle(fmt.Sprintf("Total: %d", cfg.MobsSpawnMax))
		cfg.Save()
	})

	mMobs.AddSubMenuItem("Total--", "").Click(func() {
		cfg.MobsSpawnMax--
		mMobsMaximo.SetTitle(fmt.Sprintf("Total: %d", cfg.MobsSpawnMax))
		cfg.Save()
	})

	mMobs.AddSubMenuItem("Spawn", "Spawn").Click(func() {
		SpawnRandom(1)
	})

	mMobsCycle := mMobs.AddSubMenuItem("Enable cycle", "Enable cycle")
	mMobsCycle.Click(func() {
		cfg.MobsCycle = !cfg.MobsCycle
		if cfg.MobsCycle {
			mMobsCycle.Check()
		} else {
			mMobsCycle.Uncheck()
		}
		cfg.Save()
	})
	if cfg.MobsCycle {
		mMobsCycle.Check()
	}

	systray.AddSeparator()

	systray.AddMenuItem("Reload", "Reload").Click(func() {
		cfg.Load()
		// was stopping the program
		// setScreenArea()
		loadMobsConfig()
		loadedGifs = nil
		loadedGifs = make(map[string][]*ebiten.Image)
		nextSpawn = 1
	})

	/**
	* ========== END ==========
	 */
	systray.AddSeparator()

	systray.AddMenuItem("Close", "Close").Click(func() {
		cfg.Save()
		systray.Quit()
	})
}

func onExit() {}

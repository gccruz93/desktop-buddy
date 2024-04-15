package main

import (
	"desktop-buddy/assets"
	"desktop-buddy/internal/core"
	"desktop-buddy/internal/emotes"
	"desktop-buddy/internal/mobs"
	"fmt"

	"github.com/energye/systray"
)

func onReady() {
	systray.SetIcon(assets.Icontray)
	systray.SetTitle(core.Title)
	systray.SetTooltip(core.Title)
	systray.SetOnRClick(func(menu systray.IMenu) {
		menu.ShowMenu()
	})

	systray.AddMenuItem("v0.2.0 - @twpax", "v0.2.0 - @twpax").Disable()

	systray.AddMenuItem("Close", "Close").Click(func() {
		core.Cfg.Save()
		systray.Quit()
	})

	systray.AddSeparator()

	mTaskbar := systray.AddMenuItem("Hide from taskbar", "Hide from taskbar")
	mTaskbar.Click(func() {
		core.Cfg.SkipTaskbar = !core.Cfg.SkipTaskbar
		if core.Cfg.SkipTaskbar {
			mTaskbar.Check()
		} else {
			mTaskbar.Uncheck()
		}
		core.Cfg.Save()
	})
	if core.Cfg.SkipTaskbar {
		mTaskbar.Check()
	}

	systray.AddMenuItem("Reload", "Reload").Click(func() {
		core.Cfg.Load()
		mobs.LoadConfig()
		emotes.LoadConfig()
	})

	systray.AddSeparator()

	/**
	* ========== MOBS ==========
	 */
	mMobs := systray.AddMenuItem("Mobs", "Mobs")
	mMobsSpawnTotal := mMobs.AddSubMenuItem(fmt.Sprintf("Total: %d", core.Cfg.MobsSpawnTotal), "")
	mMobsSpawnTotal.Disable()

	mMobs.AddSubMenuItem("Total++", "").Click(func() {
		core.Cfg.MobsSpawnTotal++
		mMobsSpawnTotal.SetTitle(fmt.Sprintf("Total: %d", core.Cfg.MobsSpawnTotal))
		core.Cfg.Save()
	})

	mMobs.AddSubMenuItem("Total--", "").Click(func() {
		core.Cfg.MobsSpawnTotal--
		mMobsSpawnTotal.SetTitle(fmt.Sprintf("Total: %d", core.Cfg.MobsSpawnTotal))
		core.Cfg.Save()
	})

	mMobsSpawnCycle := mMobs.AddSubMenuItem("Enable cycle", "Enable cycle")
	mMobsSpawnCycle.Click(func() {
		core.Cfg.MobsSpawnCycle = !core.Cfg.MobsSpawnCycle
		if core.Cfg.MobsSpawnCycle {
			mMobsSpawnCycle.Check()
		} else {
			mMobsSpawnCycle.Uncheck()
		}
		core.Cfg.Save()
	})
	if core.Cfg.MobsSpawnCycle {
		mMobsSpawnCycle.Check()
	}

	systray.AddMenuItem("Spawn", "Spawn").Click(func() {
		mobs.SpawnRandom(1)
	})
}

func onExit() {}

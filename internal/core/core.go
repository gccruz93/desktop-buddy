package core

import (
	"desktop-buddy/internal/config"
)

var (
	Title         = "Desktop Buddy"
	ScreenHeight  = 0.0
	ScreenWidth   = 0.0
	FrameTick     = 0
	NextSpawnTick = 1
)

var (
	Cfg config.Config
)

// const sampleRate = 48000

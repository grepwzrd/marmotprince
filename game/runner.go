package game

import (
	"github.com/grepwzrd/marmotprince/locations"
	"github.com/grepwzrd/marmotprince/player"
	"github.com/grepwzrd/marmotprince/ui"
)

func Start() {
	ui.ClearScreen()
	ui.PrintWelcome()
	prince := player.InitializePlayer()
	world := locations.BuildWorldMap()
	ui.DisplayPlayerStatus(prince, world)
}

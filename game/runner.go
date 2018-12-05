package game

import (
	"fmt"
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
	// ui.displaycurrentchoices based on current world, merchant, and game state?
	// get user input
	// interpret choice
	// do choice and tick appropriate amount
	choice := ui.GetUserInput()
}

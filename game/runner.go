package game

import (
	"fmt"
	"github.com/grepwzrd/marmotprince/locations"
	"github.com/grepwzrd/marmotprince/player"
	"github.com/grepwzrd/marmotprince/rules"
	"github.com/grepwzrd/marmotprince/ui"
)

func Start() {
	ui.ClearScreen()
	ui.PrintWelcome()
	prince := player.InitializePlayer()
	world := locations.BuildWorldMap()
	ui.DisplayPlayerStatus(prince, world)
	// interpret choice
	// do choice and tick appropriate amount
	choices := rules.GenerateCurrentChoices(prince, world)
	ui.DisplayChoices(choices)
	choice := ui.GetUserInput()
	fmt.Println(choice)
}

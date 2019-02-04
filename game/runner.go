package game

import (
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
	gameloop(prince, world)
}

func gameloop(prince player.Merchant, world locations.WorldMap) {
	for {
		ui.DisplayPlayerStatus(prince, world)
		choices := rules.GenerateCurrentChoices(prince, world)
		ui.DisplayChoices(choices)
		choice := ui.GetUserInput()
		if rules.ValidChoice(choice, choices) {
			//chosen := rules.Chosen(choice, choices)
			world = world.Tick()
		} else {
			ui.DisplayInputError()
		}
	}
}

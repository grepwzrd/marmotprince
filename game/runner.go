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
	gameloop(prince, world)
}

func gameloop(prince player.Merchant, world locations.WorldMap) {
	for {
		ui.DisplayPlayerStatus(prince, world)
		choices := rules.GenerateCurrentChoices(prince, world)
		ui.DisplayChoices(choices)
		choice := ui.GetUserInput()
		if rules.ValidChoice(choice, choices) { // check for conventions and idiomatic way to do all this stuff
			chosen := rules.Chosen(choice, choices)
			tick(chosen, prince, world)
		} else {
			ui.DisplayInputError()
		}
	}
}

func tick(choice rules.Choice, prince player.Merchant, world locations.WorldMap) {
	// perform updates based on the data in the choice
	fmt.Println("doing stuff bc you chose", choice.Keyword)
}

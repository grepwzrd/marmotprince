package ui

import (
	"fmt"
	"github.com/grepwzrd/marmotprince/locations"
	"github.com/grepwzrd/marmotprince/player"
	"github.com/grepwzrd/marmotprince/rules"
	"os"
	"os/exec"
)

// change this to use bufio

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func PrintWelcome() {
	var message string = `
		*---*---*---*---*---*---*---*
		|                           |
		|            @              |
		|           /_\             |
		|          x-x-x            |
		|        (d o_o b)          |
		|          \ . /            |
		|           |v|             |
		|                           |
		| Welcome to Marmot Prince! |
		|                           |
		*---*---*---*---*---*---*---*
		`
	fmt.Println(message)
}

func displayPlayerInformation(name string, gold int) {
	info := fmt.Sprintf("%s, %d gold coins", name, gold)
	fmt.Println(info)
}

func displayPlayerLocation(posx int, posy int, cities []locations.City) {
	playerCoordinates := [2]int{posx, posy}
	if cities[0].Coordinates() == playerCoordinates {
		fmt.Println(fmt.Sprintf("Location: %s", cities[0].Name))
	} else if cities[1].Coordinates() == playerCoordinates {
		fmt.Println(fmt.Sprintf("Location: %s", cities[1].Name))
	} else if cities[2].Coordinates() == playerCoordinates {
		fmt.Println(fmt.Sprintf("Location: %s", cities[2].Name))
	} else {
		fmt.Println("Location: Wilderness")
	}
}

func displayMap(w locations.WorldMap) {
	const Border = "⬜"
	const Grass = "🌱"
	const Player = "🐹"
}

func DisplayPlayerStatus(prince player.Merchant, world locations.WorldMap) {
	displayMap(world)
	displayPlayerInformation(prince.Name, prince.Gold)
	displayPlayerLocation(prince.Positionx, prince.Positiony, world.Cities)
}

func DisplayInputError() {
	fmt.Println("!!! Error. That does not appear to be a valid choice.")
}

func DisplayChoices(choices []rules.Choice) {
	header := "--- Please Select a Choice ---"
	fmt.Println(header)
	for i := 0; i < len(choices); i++ {
		fmt.Println(choices[i].Keyword)
	}
}

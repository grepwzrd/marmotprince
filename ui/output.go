package ui

import (
	"fmt"
	"github.com/grepwzrd/marmotprince/locations"
	"github.com/grepwzrd/marmotprince/player"
	"os"
	"os/exec"
)

func ClearScreen() {
	// check out fmt.Scanln() for taking user input from terminal
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
		|         ( o_o )           |
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

func DisplayPlayerStatus(prince player.Merchant, world locations.WorldMap) {
	displayPlayerInformation(prince.Name, prince.Gold)
	displayPlayerLocation(prince.Positionx, prince.Positiony, world.Cities)
}

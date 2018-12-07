package rules

import (
	"github.com/grepwzrd/marmotprince/locations"
	"github.com/grepwzrd/marmotprince/player"
)

// this might be a place for
// methods and interfaces
type Choice struct {
	Keyword string
}

func GenerateCurrentChoices(prince player.Merchant, world locations.WorldMap) []Choice {
	choices := []Choice{}
	choices = append(choices, Choice{Keyword: "rest"})
	choices = append(choices, Choice{Keyword: "travel"})
	choices = append(choices, Choice{Keyword: "shop"})
	return choices
}

func ValidChoice(choice string, choices []Choice) bool {
	// if choices has a keyword that matches choice
	return true
}

func Chosen(choice string, choices []Choice) Choice {
	//
	return choices[0]
}

//func PerformChoice(choice Choice, prince Merchant, world WorldMap) {
//}

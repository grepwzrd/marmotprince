package ui

import (
	"fmt"
	"os"
	"os/exec"
)

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

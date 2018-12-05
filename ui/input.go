package ui

import (
	"bufio"
	"os"
)

func GetUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

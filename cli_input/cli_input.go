package cli_input

import (
	"fmt"
	"io"
)

type Parser interface {
	PlayerName() string
	Health() int
	Strength() int
	Attack() int
}

type input struct {
	playerName string
	health     int
	strength   int
	attack     int
}

func (i input) PlayerName() string {
	return i.playerName
}

func (i input) Health() int {
	return i.health
}

func (i input) Strength() int {
	return i.strength
}

func (i input) Attack() int {
	return i.attack
}

func CliInput(reader io.Reader) (cliInput input) {
	fmt.Fscanf(reader, "%s %d %d %d", &cliInput.playerName, &cliInput.health, &cliInput.strength, &cliInput.attack)
	return
}

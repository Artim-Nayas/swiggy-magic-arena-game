package models

import (
	"fmt"
	"os"

)

type Game struct {
	PlayerA *Player
	PlayerB *Player
	active  bool
}

func NewGame(playerA, playerB *Player) *Game {
	return &Game{
		PlayerA: playerA,
		PlayerB: playerB,
		active:  true,
	}
}

func (g *Game) IsActive() bool {
	return g.active
}

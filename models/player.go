package models

import (
	"fmt"

	game_action "github.com/ArtimNayas/magic_arena_game/game_actions"
)

type Player struct {
	Name     string
	Health   int
	Strength int
	Attack   int
}

func NewPlayer(name string, health, strength, attack int) *Player {
	return &Player{
		Name:     name,
		Health:   health,
		Strength: strength,
		Attack:   attack,
	}
}
func (p *Player) AttackDamage() int {
	return p.Attack * game_action.RollDice()
}

func (p *Player) DefendDamage() int {
	return p.Strength * game_action.RollDice()
}

func (p *Player) String() string {
	return fmt.Sprintf("%s (Health: %d, Strength: %d, Attack: %d)", p.Name, p.Health, p.Strength, p.Attack)
}

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

func (p *Player) TakeDamage(damage int) {
	p.Health -= damage
	if p.Health < 0 {
		p.Health = 0
	}
}

func (p *Player) IsAlive() bool {
	return p.Health > 0
}

func (p *Player) RollAttackDice() int {
	return game_action.RollDice()
}

func (p *Player) RollDefendDice() int {
	return game_action.RollDice()
}

func (p *Player) String() string {
	return fmt.Sprintf("%s (Health: %d, Strength: %d, Attack: %d)", p.Name, p.Health, p.Strength, p.Attack)
}

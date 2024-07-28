package models

import (
	"fmt"

	"github.com/ArtimNayas/magic_arena_game/cli_output"
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

func (g *Game) GetWinner() *Player {
	if !g.PlayerA.IsAlive() && g.PlayerB.IsAlive() {
		return g.PlayerB
	} else if g.PlayerA.IsAlive() && !g.PlayerB.IsAlive() {
		return g.PlayerA
	}
	return nil
}

func (g *Game) Attack(attacker, defender *Player) {
	damage := attacker.AttackDamage()
	defended := defender.DefendDamage()
	finalDamage := damage - defended
	if finalDamage < 0 {
		finalDamage = 0
	}
	defender.TakeDamage(finalDamage)
	cli_output.Render(fmt.Sprintf("%s attacks %s: attack damage = %d, defended damage = %d, final damage = %d",
		attacker.Name, defender.Name, damage, defended, finalDamage))

	if !defender.IsAlive() {
		g.active = false
	}
}

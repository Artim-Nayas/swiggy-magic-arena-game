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

func (g *Game) PlayTurn() {
	attacker, defender := g.selectAttackerAndDefender()
	attackerDiceRoll := attacker.RollAttackDice()
	defenderDiceRoll := defender.RollDefendDice()

	damage := attacker.Attack * attackerDiceRoll
	defended := defender.Strength * defenderDiceRoll
	finalDamage := damage - defended
	if finalDamage < 0 {
		finalDamage = 0
	}
	defender.TakeDamage(finalDamage)

	cli_output.Render(fmt.Sprintf("%s attacks %s: attack roll = %d, attack damage = %d, defend roll = %d, defended damage = %d, final damage = %d",
		attacker.Name, defender.Name, attackerDiceRoll, damage, defenderDiceRoll, defended, finalDamage))

	cli_output.Render(fmt.Sprintf("%s's remaining health: %d", defender.Name, defender.Health))
	cli_output.Render(fmt.Sprintf("%s's remaining health: %d", attacker.Name, attacker.Health))

	if !defender.IsAlive() {
		g.active = false
	}
}

func (g *Game) selectAttackerAndDefender() (*Player, *Player) {
	if g.PlayerA.Health < g.PlayerB.Health {
		return g.PlayerA, g.PlayerB
	}
	return g.PlayerB, g.PlayerA
}

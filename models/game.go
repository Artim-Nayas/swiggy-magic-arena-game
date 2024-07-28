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
	// Determine the order of attack based on health
	var firstAttacker, secondAttacker, firstDefender, secondDefender *Player

	if g.PlayerA.Health <= g.PlayerB.Health {
		firstAttacker, firstDefender = g.PlayerA, g.PlayerB
		secondAttacker, secondDefender = g.PlayerB, g.PlayerA
	} else {
		firstAttacker, firstDefender = g.PlayerB, g.PlayerA
		secondAttacker, secondDefender = g.PlayerA, g.PlayerB
	}

	// First attack
	g.attack(firstAttacker, firstDefender)

	// Check if the first defender is still alive before the second attack
	if firstDefender.IsAlive() {
		g.attack(secondAttacker, secondDefender)
	}

	// Check if the game should still be active
	if !g.PlayerA.IsAlive() || !g.PlayerB.IsAlive() {
		g.active = false
	}
}

func (g *Game) attack(attacker, defender *Player) {
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
}

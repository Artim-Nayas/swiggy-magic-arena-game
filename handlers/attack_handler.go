package handler

import (
	"github.com/ArtimNayas/magic_arena_game/models"
)

const attackOperation Operation = "attack"

func AttackHandler(game *models.Game) {
	playerA := game.PlayerA
	playerB := game.PlayerB
	if playerA.Health < playerB.Health {
		game.Attack(playerA, playerB)
	} else {
		game.Attack(playerB, playerA)
	}
}

func init() {
	RegisterHandler(attackOperation, AttackHandler)
}

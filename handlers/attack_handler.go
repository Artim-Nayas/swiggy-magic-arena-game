package handler

import (
	"github.com/ArtimNayas/magic_arena_game/models"
)

const attackOperation Operation = "attack"

func AttackHandler(game *models.Game) {
	game.PlayTurn()
}

func init() {
	RegisterHandler(attackOperation, AttackHandler)
}

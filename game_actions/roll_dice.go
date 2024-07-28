package game_action

import (
	"math/rand"
)

func RollDice() int {
	return rand.Intn(6) + 1
}

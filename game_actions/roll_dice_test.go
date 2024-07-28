package game_action

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRollDice(t *testing.T) {
	for i := 0; i < 100; i++ {
		result := RollDice()
		assert.True(t, result >= 1 && result <= 6, "RollDice() result out of range")
	}
}

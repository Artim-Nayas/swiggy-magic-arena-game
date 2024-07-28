package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	playerA := NewPlayer("PlayerA", 100, 10, 15)
	playerB := NewPlayer("PlayerB", 80, 8, 12)
	game := NewGame(playerA, playerB)

	assert.True(t, game.IsActive())
	assert.Nil(t, game.GetWinner())
}

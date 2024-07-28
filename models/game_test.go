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

func TestPlayerTakeDamageAndIsAlive(t *testing.T) {
	player := NewPlayer("PlayerA", 100, 10, 15)

	player.TakeDamage(50)
	assert.Equal(t, 50, player.Health)
	assert.True(t, player.IsAlive())

	player.TakeDamage(60)
	assert.Equal(t, 0, player.Health)
	assert.False(t, player.IsAlive())
}

func TestGameAttackAndWinner(t *testing.T) {
	playerA := NewPlayer("PlayerA", 100, 10, 15)
	playerB := NewPlayer("PlayerB", 80, 8, 12)
	game := NewGame(playerA, playerB)

	playerA.TakeDamage(100) // PlayerA is now at 0 health
	assert.False(t, playerA.IsAlive())
	assert.True(t, game.IsActive())

	game.Attack(playerB, playerA)
	assert.False(t, game.IsActive())
	assert.Equal(t, playerB, game.GetWinner())
}

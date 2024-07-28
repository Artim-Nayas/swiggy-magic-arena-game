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

func TestGamePlayTurnAndWinner(t *testing.T) {
	playerA := NewPlayer("PlayerA", 100, 10, 15)
	playerB := NewPlayer("PlayerB", 80, 8, 12)
	game := NewGame(playerA, playerB)

	// Simulate the game turns until one player wins
	for game.IsActive() {
		game.PlayTurn()
	}

	winner := game.GetWinner()
	assert.NotNil(t, winner)
	assert.True(t, winner.IsAlive())
	assert.False(t, game.IsActive())
}

func TestSelectAttackerAndDefender(t *testing.T) {
	playerA := NewPlayer("PlayerA", 50, 10, 15)
	playerB := NewPlayer("PlayerB", 80, 8, 12)
	game := NewGame(playerA, playerB)

	attacker, defender := game.selectAttackerAndDefender()
	assert.Equal(t, playerA, attacker)
	assert.Equal(t, playerB, defender)

	playerA.TakeDamage(50) // PlayerA's health is now lower than PlayerB's health
	attacker, defender = game.selectAttackerAndDefender()
	assert.Equal(t, playerA, attacker)
	assert.Equal(t, playerB, defender)

	playerB.TakeDamage(60) // PlayerB's health is now lower than PlayerA's health
	attacker, defender = game.selectAttackerAndDefender()
	assert.Equal(t, playerB, attacker)
	assert.Equal(t, playerA, defender)
}

func TestGameWithExactDamage(t *testing.T) {
	playerA := NewPlayer("PlayerA", 10, 5, 15)
	playerB := NewPlayer("PlayerB", 10, 5, 15)
	game := NewGame(playerA, playerB)

	// Ensure the game ends when a player takes exact damage to zero health
	for game.IsActive() {
		game.PlayTurn()
	}

	assert.False(t, game.IsActive())
	winner := game.GetWinner()
	assert.NotNil(t, winner)
	assert.True(t, winner.IsAlive() || playerA.Health == 0 && playerB.Health == 0)
}

func TestGameDraw(t *testing.T) {
	playerA := NewPlayer("PlayerA", 10, 5, 15)
	playerB := NewPlayer("PlayerB", 10, 5, 15)
	game := NewGame(playerA, playerB)

	// Simulate both players dying in the same turn
	playerA.TakeDamage(10)
	playerB.TakeDamage(10)
	assert.False(t, playerA.IsAlive())
	assert.False(t, playerB.IsAlive())

	game.PlayTurn() // This should end the game

	assert.False(t, game.IsActive())
	assert.Nil(t, game.GetWinner())
}

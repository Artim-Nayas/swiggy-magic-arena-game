package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerInitialization(t *testing.T) {
	player := NewPlayer("PlayerA", 100, 10, 15)
	assert.Equal(t, "PlayerA", player.Name)
	assert.Equal(t, 100, player.Health)
	assert.Equal(t, 10, player.Strength)
	assert.Equal(t, 15, player.Attack)
}

func TestPlayerTakeDamage(t *testing.T) {
	player := NewPlayer("PlayerA", 100, 10, 15)
	player.TakeDamage(20)
	assert.Equal(t, 80, player.Health)
	assert.True(t, player.IsAlive())
}

func TestPlayerTakeFatalDamage(t *testing.T) {
	player := NewPlayer("PlayerA", 100, 10, 15)
	player.TakeDamage(100)
	assert.Equal(t, 0, player.Health)
	assert.False(t, player.IsAlive())
}

func TestPlayerRollAttackDice(t *testing.T) {
	player := NewPlayer("PlayerA", 100, 10, 15)
	for i := 0; i < 100; i++ { // Testing multiple rolls for randomness
		attackDice := player.RollAttackDice()
		assert.True(t, attackDice >= 1 && attackDice <= 6)
	}
}

func TestPlayerRollDefendDice(t *testing.T) {
	player := NewPlayer("PlayerA", 100, 10, 15)
	for i := 0; i < 100; i++ { // Testing multiple rolls for randomness
		defendDice := player.RollDefendDice()
		assert.True(t, defendDice >= 1 && defendDice <= 6)
	}
}

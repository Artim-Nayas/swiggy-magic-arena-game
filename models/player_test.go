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
func TestPlayerAttackDamage(t *testing.T) {
	player := NewPlayer("PlayerA", 100, 10, 15)
	attackDamage := player.AttackDamage()
	assert.True(t, attackDamage >= 15 && attackDamage <= 90) // 1 to 6 times 15
}

func TestPlayerDefendDamage(t *testing.T) {
	player := NewPlayer("PlayerA", 100, 10, 15)
	defendDamage := player.DefendDamage()
	assert.True(t, defendDamage >= 10 && defendDamage <= 60) // 1 to 6 times 10
}

package handler

import (
	"io"
	"os"
	"testing"

	"github.com/ArtimNayas/magic_arena_game/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAttackHandler(t *testing.T) {
	playerA := models.NewPlayer("PlayerA", 100, 10, 15)
	playerB := models.NewPlayer("PlayerB", 80, 8, 12)
	game := models.NewGame(playerA, playerB)

	handlerFunc, exists := GetHandler(attackOperation)
	require.True(t, exists, "Attack handler should exist")
	require.NotNil(t, handlerFunc, "Attack handler should not be nil")

	handlerFunc(game)

	assert.True(t, game.IsActive(), "Game should still be active after the attack")
	assert.True(t, playerA.IsAlive(), "Player A should still be alive after the attack")
	assert.True(t, playerB.IsAlive(), "Player B should still be alive after the attack")
}

func TestAttackHandlerOutput(t *testing.T) {
	defer setupStdout()()
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)

	os.Stdout = fakeStdout
	playerA := models.NewPlayer("PlayerA", 100, 10, 15)
	playerB := models.NewPlayer("PlayerB", 80, 8, 12)
	game := models.NewGame(playerA, playerB)

	handlerFunc, exists := GetHandler(attackOperation)
	require.True(t, exists, "Attack handler should exist")
	require.NotNil(t, handlerFunc, "Attack handler should not be nil")

	handlerFunc(game)

	fakeStdout.Close()
	bytes, err := io.ReadAll(r)
	require.NoError(t, err)
	r.Close()

	expectedOutput := "PlayerB attacks PlayerA: attack damage = 72, defended damage = 40, final damage = 32\n"
	assert.Equal(t, expectedOutput, string(bytes))

	r, fakeStdout, err = os.Pipe()
	require.NoError(t, err)
	os.Stdout = fakeStdout

	handlerFunc(game)

	fakeStdout.Close()
	bytes, err = io.ReadAll(r)
	require.NoError(t, err)
	r.Close()

	expectedOutput = "PlayerA attacks PlayerB: attack damage = 90, defended damage = 48, final damage = 42\n"
	assert.Equal(t, expectedOutput, string(bytes))
}

func TestAttackHandlerRegistration(t *testing.T) {
	handlerFunc, exists := GetHandler(attackOperation)
	assert.True(t, exists, "Attack handler should exist")
	assert.NotNil(t, handlerFunc, "Attack handler should not be nil")
	assert.IsType(t, func(game *models.Game) {}, handlerFunc)
}

func setupStdout() func() {
	originalStdout := os.Stdout
	return func() { os.Stdout = originalStdout }
}

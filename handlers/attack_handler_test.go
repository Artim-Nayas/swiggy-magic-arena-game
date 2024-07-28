package handler

import (
	"bytes"
	"fmt"
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

	handlerFunc, exists := GetHandler("attack")
	require.True(t, exists, "Attack handler should exist")
	require.NotNil(t, handlerFunc, "Attack handler should not be nil")

	handlerFunc(game)

	assert.True(t, game.IsActive(), "Game should still be active after the attack")
	assert.True(t, playerA.IsAlive(), "Player A should still be alive after the attack")
	assert.True(t, playerB.IsAlive(), "Player B should still be alive after the attack")
}

func TestAttackHandlerOutput(t *testing.T) {
	// Capture the output
	r, w, err := os.Pipe()
	require.NoError(t, err)

	originalStdout := os.Stdout
	os.Stdout = w

	playerA := models.NewPlayer("PlayerA", 100, 10, 15)
	playerB := models.NewPlayer("PlayerB", 80, 8, 12)
	game := models.NewGame(playerA, playerB)

	handlerFunc, exists := GetHandler("attack")
	require.True(t, exists, "Attack handler should exist")
	require.NotNil(t, handlerFunc, "Attack handler should not be nil")

	handlerFunc(game)

	// Restore original stdout
	w.Close()
	os.Stdout = originalStdout

	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	require.NoError(t, err)
	output := buf.String()

	// The output is dynamic due to random dice rolls, so we need to check for expected substrings
	expectedSubstrings := []string{
		"attacks",
		"attack roll",
		"defend roll",
		"attack damage",
		"defended damage",
		"final damage",
		"remaining health",
	}

	for _, substring := range expectedSubstrings {
		assert.Contains(t, output, substring, fmt.Sprintf("Output should contain %s", substring))
	}
}

func TestAttackHandlerRegistration(t *testing.T) {
	handlerFunc, exists := GetHandler("attack")
	assert.True(t, exists, "Attack handler should exist")
	assert.NotNil(t, handlerFunc, "Attack handler should not be nil")
	assert.IsType(t, func(game *models.Game) {}, handlerFunc)
}

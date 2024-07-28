package manager

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStartGame(t *testing.T) {
	// Simulate user input
	input := "PlayerA 100 10 15\nPlayerB 80 8 12\n"
	rStdin, wStdin, err := os.Pipe()
	require.NoError(t, err)

	_, err = wStdin.WriteString(input)
	require.NoError(t, err)
	wStdin.Close()

	originalStdin := os.Stdin
	os.Stdin = rStdin
	defer func() { os.Stdin = originalStdin }()

	// Capture output
	rStdout, wStdout, err := os.Pipe()
	require.NoError(t, err)

	originalStdout := os.Stdout
	os.Stdout = wStdout
	defer func() { os.Stdout = originalStdout }()

	var buf bytes.Buffer
	done := make(chan struct{})
	go func() {
		io.Copy(&buf, rStdout)
		close(done)
	}()

	// Start the game
	StartGame()

	wStdout.Close()
	<-done

	// Check the output
	output := buf.String()
	expectedOutputs := []string{
		"Enter details for Player A:",
		"Enter details for Player B:",
	}

	for _, expected := range expectedOutputs {
		assert.Contains(t, output, expected)
	}
}

func TestCreatePlayer(t *testing.T) {
	input := "PlayerA 100 10 15\n"
	rStdin, wStdin, err := os.Pipe()
	require.NoError(t, err)

	_, err = wStdin.WriteString(input)
	require.NoError(t, err)
	wStdin.Close()

	originalStdin := os.Stdin
	os.Stdin = rStdin
	defer func() { os.Stdin = originalStdin }()

	player := createPlayer("Enter details for Player A:")
	assert.Equal(t, "PlayerA", player.Name)
	assert.Equal(t, 100, player.Health)
	assert.Equal(t, 10, player.Strength)
	assert.Equal(t, 15, player.Attack)
}

func TestGameManagerIntegration(t *testing.T) {
	input := "PlayerA 100 10 15\nPlayerB 80 8 12\nattack\n"
	rStdin, wStdin, err := os.Pipe()
	require.NoError(t, err)

	_, err = wStdin.WriteString(input)
	require.NoError(t, err)
	wStdin.Close()

	originalStdin := os.Stdin
	os.Stdin = rStdin
	defer func() { os.Stdin = originalStdin }()

	rStdout, wStdout, err := os.Pipe()
	require.NoError(t, err)

	originalStdout := os.Stdout
	os.Stdout = wStdout
	defer func() { os.Stdout = originalStdout }()

	var buf bytes.Buffer
	done := make(chan struct{})
	go func() {
		io.Copy(&buf, rStdout)
		close(done)
	}()

	StartGame()

	wStdout.Close()
	<-done

	output := buf.String()
	assert.Contains(t, output, "Enter details for Player A:")
	assert.Contains(t, output, "Enter details for Player B:")
}

package cli_output

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRender(t *testing.T) {
	defer setupStdout()()
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)

	os.Stdout = fakeStdout
	Render("output should be this")

	fakeStdout.Close()
	bytes, err := io.ReadAll(r)
	require.NoError(t, err)
	r.Close()
	assert.Equal(t, "output should be this\n", string(bytes))
}

func TestRenderInvalidOperation(t *testing.T) {
	defer setupStdout()()
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)

	os.Stdout = fakeStdout
	RenderInvalidOperation()

	fakeStdout.Close()
	bytes, err := io.ReadAll(r)
	require.NoError(t, err)
	r.Close()
	assert.Equal(t, "Invalid operation\n", string(bytes))
}

func TestRenderCliInput(t *testing.T) {
	defer setupStdout()()
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)

	os.Stdout = fakeStdout
	RenderCliInput()

	fakeStdout.Close()
	bytes, err := io.ReadAll(r)
	require.NoError(t, err)
	r.Close()
	assert.Equal(t, "\n>", string(bytes))
}

func setupStdout() func() {
	originalStdout := os.Stdout
	return func() { os.Stdout = originalStdout }
}

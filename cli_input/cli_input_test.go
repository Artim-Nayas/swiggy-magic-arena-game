package cli_input

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCliInput(t *testing.T) {
	t.Run("should accept player name and numerical attributes", func(t *testing.T) {
		var reader io.Reader = strings.NewReader("PlayerA 100 10 15")
		input := CliInput(reader)
		assert.Equal(t, "PlayerA", input.PlayerName())
		assert.Equal(t, 100, input.Health())
		assert.Equal(t, 10, input.Strength())
		assert.Equal(t, 15, input.Attack())
	})
}

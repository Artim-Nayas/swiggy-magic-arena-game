package handler

import (
	"testing"

	"github.com/ArtimNayas/magic_arena_game/models"
	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler(t *testing.T) {
	t.Run("should be able to register handlers with operation and handler func", func(t *testing.T) {
		defer resetHandler()()
		assert.NotPanics(t, func() {
			RegisterHandler(Operation("attack"), func(game *models.Game) {})
		})
	})

	t.Run("should panic if Operation already exists", func(t *testing.T) {
		defer resetHandler()()
		RegisterHandler(Operation("attack"), func(game *models.Game) {})

		assert.Panics(t, func() {
			RegisterHandler(Operation("attack"), func(game *models.Game) {})
		})
	})
}

func resetHandler() func() {
	initialHandler := handlers
	handlers = make(map[Operation]func(*models.Game))
	return func() { handlers = initialHandler }
}

func TestGetHandler(t *testing.T) {
	t.Run("should return a handler given an Operation", func(t *testing.T) {
		defer resetHandler()()
		RegisterHandler(Operation("attack"), func(game *models.Game) {})

		h, exists := GetHandler(Operation("attack"))
		assert.True(t, exists, "Handler should exist")
		assert.NotNil(t, h, "Handler should not be nil")
	})

	t.Run("should return nil when Operation is not found", func(t *testing.T) {
		defer resetHandler()()

		h, exists := GetHandler(Operation("nonExistentOperation"))
		assert.False(t, exists, "Handler should not exist")
		assert.Nil(t, h, "Handler should be nil")
	})
}

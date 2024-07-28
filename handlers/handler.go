package handler

import (
	"fmt"

	"github.com/ArtimNayas/magic_arena_game/models"
)

type Operation string

var handlers = make(map[Operation]func(*models.Game))

func RegisterHandler(op Operation, handler func(*models.Game)) {
	if _, found := handlers[op]; found {
		err := fmt.Errorf("duplicate handlers registration for Operation: %s", op)
		panic(err)
	}
	handlers[op] = handler
}

func GetHandler(op Operation) (func(*models.Game), bool) {
	handler, exists := handlers[op]
	return handler, exists
}

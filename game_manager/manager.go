package manager

import (
	"fmt"

	"github.com/ArtimNayas/magic_arena_game/models"
)

func StartGame() {
	playerA := createPlayer("Enter details for Player A:")
	playerB := createPlayer("Enter details for Player B:")

	game := models.NewGame(playerA, playerB)

	for game.IsActive() {
		game.Attack(playerA, playerB)
		if !game.IsActive() {
			break
		}
		game.Attack(playerB, playerA)
	}

	winner := game.GetWinner()
	if winner != nil {
		fmt.Printf("%s wins the game!\n", winner.Name)
	}
}

func createPlayer(prompt string) *models.Player {
	var name string
	var health, strength, attack int
	fmt.Println(prompt)
	fmt.Scanf("%s %d %d %d", &name, &health, &strength, &attack)
	return models.NewPlayer(name, health, strength, attack)
}

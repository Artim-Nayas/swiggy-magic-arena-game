package manager

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ArtimNayas/magic_arena_game/models"
)

func StartGame() {
	reader := bufio.NewReader(os.Stdin)
	for {
		playerA := createPlayer(reader, "Enter details for Player A (name health strength attack):")
		playerB := createPlayer(reader, "Enter details for Player B (name health strength attack):")

		game := models.NewGame(playerA, playerB)

		for game.IsActive() {
			game.PlayTurn()
		}

		winner := game.GetWinner()
		if winner != nil {
			fmt.Printf("%s wins the game!\n", winner.Name)
		} else {
			fmt.Println("It's a draw!")
		}

		fmt.Println("Do you want to play again? (yes/no)")
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(response)
		if strings.ToLower(response) != "yes" {
			break
		}
	}
}

func createPlayer(reader *bufio.Reader, prompt string) *models.Player {
	fmt.Println(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var name string
	var health, strength, attack int
	fmt.Sscanf(input, "%s %d %d %d", &name, &health, &strength, &attack)

	return models.NewPlayer(name, health, strength, attack)
}

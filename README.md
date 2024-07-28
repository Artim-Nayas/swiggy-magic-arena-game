# Magic Arena Game

Welcome to the Magic Arena Game! This is a command-line game where two players with different attributes battle against each other using dice rolls. The game allows players to attack and defend, and the outcome of each round is determined by the roll of the dice.

## Instructions

### Prerequisites

- Go 1.16 or later

### Build the Game

To build the game, run the following command in the root directory of the repository:

```sh
go build -o magic_arena_game main.go
```

### Run the Game

To start the game, execute the following command:

```sh
./magic_arena_game
```

### Running Tests

The project uses the `testify` package for testing. To run the tests, use the following command:

```sh
go test ./...
```

### Game Instructions

1. **Player Setup**:

   - The game will prompt you to enter details for Player A and Player B.
   - Input the name, health, strength, and attack values for each player in the format: `name health strength attack`.

2. **Gameplay**:

   - Players take turns attacking each other.
   - During each attack, the attacking player rolls a die to determine the attack damage, and the defending player rolls a die to determine the defended damage.
   - The damage dealt to the defending player is calculated as: `final damage = (attack damage) - (defended damage)`.
   - After each attack, the remaining health of each player is displayed.

3. **Winning the Game**:

   - The game continues until one player's health drops to 0 or below.
   - The player with remaining health is declared the winner.

4. **Play Again**:
   - After a game ends, you will be asked if you want to play again.
   - Enter `yes` to start a new game or `no` to exit.

### Example Game Flow

```
Enter details for Player A (name health strength attack):
PlayerA 100 10 15

Enter details for Player B (name health strength attack):
PlayerB 80 8 12

PlayerB attacks PlayerA: attack roll = 4, attack damage = 48, defend roll = 3, defended damage = 30, final damage = 18
PlayerA's remaining health: 82
PlayerB's remaining health: 80

PlayerA attacks PlayerB: attack roll = 5, attack damage = 75, defend roll = 2, defended damage = 16, final damage = 59
PlayerA's remaining health: 82
PlayerB's remaining health: 21

...

PlayerB wins the game!
Do you want to play again? (yes/no)
```

## Project Structure

```
magic_arena_game/
├── cli_output/
│   └── cli_output.go           # Handles command-line output rendering
│   └── cli_output_test.go      # Test file for cli output
├── game_action/
│   └── roll_dice.go            # Contains logic for rolling a dice
│   └── roll_dice_test.go       # Test file for role dice
├── game_manager/
│   ├── manager.go              # Manages the game flow and player creation
│   ├── manager.go              # Manages the game flow and player creation
├── handlers/
│   ├── handler.go              # Registers and retrieves game operation handlers
│   ├── handler_test.go         # Registers and retrieves game operation handlers
│   └── attack_handler.go       # Contains logic for the attack operation
│   └── attack_handler.go       # Contains logic for the attack operation
├── models/
│   ├── game.go                 # Represents the game logic
│   ├── game_test.go            # Test file for game model
│   └── player.go               # Represents the player model and its methods
│   └── player_test.go          # Test file for player model
├── main.go                     # Entry point of the application
└── README.md                   # Project documentation
```

## Testing

To ensure the code works correctly, unit tests are provided. Run all tests using the following command:

```sh
go test ./...
```

### Contribution

Feel free to fork the repository and submit pull requests for any improvements or bug fixes.

### License

This project is licensed under the MIT License. See the `LICENSE` file for more details.

---

Enjoy playing the Magic Arena Game!

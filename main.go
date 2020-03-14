package main

import (
	"fmt"
	"strings"
)

type player struct {
	Name      string
	Character string
}

var players [3]player

func main() {

	var gameState [3][3]int
	// var gameState = [3][3]int{{0, 1, 2}, {1, 1, 1}, {0, 2, 2}}

	players[0] = player{"Empty", "_ "}
	players[1] = player{getPlayerName(1), "X "}
	players[2] = player{getPlayerName(2), "O "}

	fmt.Printf("%v you are X and %v you are O. \n", players[1].Name, players[2].Name)
	fmt.Printf("\n")
	fmt.Println("Use 3x3 matrix addreses. First line second column is 12.")

	turn := 0
	for turn < 9 {

		entry := getPlayerInput(turn, players)

		entrySlice := strings.Split(entry, "")

		if !isValidPlayerInput(entrySlice) { // test made
			continue
		}

		row, column, valid := isValidCoordinate(entrySlice, gameState)
		if !valid { // test made
			continue
		}

		gameState = applyPlayerInput(row, column, turn, gameState) // test made

		printGameState(gameState, players) // Only print in the console

		winner := getWinner(gameState, turn)
		if winner == -1 {
			turn = turn + 1
			continue
		} else {
			fmt.Printf("%s is the winner!!!", players[winner].Name)
			break
		}

	}
}

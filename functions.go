package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getPlayerName(order int) string {
	var name string

	consoleReader := bufio.NewReader(os.Stdin)
	for name == "" {
		fmt.Printf("Enter your name player Nr.%v: ", order)
		name, _ = consoleReader.ReadString('\n')
		name = fmt.Sprintf("%v", strings.TrimSpace(name))
	}

	return name
}

func getPlayerInput(turn int, p [3]player) string {
	var entry string
	if turn%2 == 0 {
		fmt.Printf("\n")
		fmt.Printf("%s your entry: ", p[1].Name)
	} else {
		fmt.Printf("\n")
		fmt.Printf("%s your entry: ", p[2].Name)
	}
	fmt.Scan(&entry)
	return entry
}

func isValidPlayerInput(str []string) bool {
	var valid = true
	if len(str) != 2 {
		fmt.Printf("The entry is not valid.")
		valid = false
	} else {
		rows := map[string]bool{"1": true, "2": true, "3": true}
		columns := map[string]bool{"1": true, "2": true, "3": true}

		if !rows[str[0]] || !columns[str[1]] {
			fmt.Printf("The entry is not valid. ")
			valid = false
		}
	}
	return valid
}

func isValidCoordinate(str []string, gameState [3][3]int) (int, int, bool) {
	var valid bool
	row, _ := strconv.Atoi(str[0])
	column, _ := strconv.Atoi(str[1])
	if gameState[row-1][column-1] != 0 {
		fmt.Println("This cell is already taken!")
		valid = false
	} else {
		valid = true
	}
	return row, column, valid
}

func applyPlayerInput(row, column, turn int, gameState [3][3]int) [3][3]int {
	var pValue *int = &gameState[row-1][column-1]
	if turn%2 == 0 {
		*pValue = 1
	} else {
		*pValue = 2
	}
	return gameState
}

// printing is fucked
func printGameState(gameState [3][3]int, p [3]player) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if gameState[i][j] == 0 {
				fmt.Printf("%v", p[0].Character)
			} else if gameState[i][j] == 1 {
				fmt.Printf("%v", p[1].Character)
			} else if gameState[i][j] == 2 {
				fmt.Printf("%v", p[2].Character)
			} else {
				fmt.Printf("Should never print this!")
				//  TODO: error
			}
		}
		fmt.Printf("\n")
	}
}

func getWinner(gameState [3][3]int, turn int) int {
	var winComb = [][]int{
		{0, 0, 0, 1, 0, 2}, // 00 01 02
		{1, 0, 1, 1, 1, 2}, // 10 11 12
		{2, 0, 2, 1, 2, 2}, // 20 21 22
		{0, 0, 1, 0, 2, 0}, // 00 10 20
		{0, 1, 1, 1, 2, 1}, // 01 11 21
		{0, 2, 1, 2, 2, 2}, // 02 12 22
		{0, 0, 1, 1, 2, 2}, // 00 11 22
		{2, 0, 1, 1, 0, 2}, // 20 11 02
	}

	for i := 0; i < 8; i++ {
		j := winComb[i]
		v1 := gameState[j[0]][j[1]]
		v2 := gameState[j[2]][j[3]]
		v3 := gameState[j[4]][j[5]]
		if v1 == v2 && v1 == v3 && v1 != 0 {
			if turn%2 == 0 {
				return 1
			}
			return 2
		}
	}
	return -1
}

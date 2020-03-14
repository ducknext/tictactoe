package main

import (
	"fmt"
	"testing"
)

func TestIsValidPlayerInput(t *testing.T) {
	t.Run("Entry 33 ", testIsValidPlayerInput([]string{"3", "3"}, true))
	t.Run("Entry 123 ", testIsValidPlayerInput([]string{"1", "2", "3"}, false))
	t.Run("Entry 41", testIsValidPlayerInput([]string{"4", "1"}, false))
	t.Run("Entry 2", testIsValidPlayerInput([]string{"2"}, false))
	t.Run("Entry ' '", testIsValidPlayerInput([]string{""}, false))
	t.Run("Entry ab", testIsValidPlayerInput([]string{"a", "b"}, false))
}

func testIsValidPlayerInput(str []string, expected bool) func(*testing.T) {
	return func(t *testing.T) {
		actual := isValidPlayerInput(str)
		if actual != expected {
			t.Error(fmt.Sprintf("With %s expected %v got %v. ", str, expected, actual))
		}
	}
}

func TestIsCoordinate(t *testing.T) {
	t.Run("Entry 11 cell has 0", testIsValidCoordinate([]string{"1", "1"}, [3][3]int{{0, 1, 2}, {1, 1, 1}, {0, 2, 2}}, true))
	t.Run("Entry 23 cell has 1", testIsValidCoordinate([]string{"2", "3"}, [3][3]int{{0, 1, 2}, {1, 1, 1}, {0, 2, 2}}, false))
}

func testIsValidCoordinate(str []string, gameState [3][3]int, expected bool) func(*testing.T) {
	return func(t *testing.T) {
		_, _, actual := isValidCoordinate(str, gameState)
		if actual != expected {
			t.Error(fmt.Sprintf("With %v expected %v got %v. ", str, expected, actual))
		}
	}
}

func TestApplyPlayerInput(t *testing.T) {
	t.Run("PlayerA cell 12 k=0", testApplyPlayerInput(1, 2, 0,
		[3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		[3][3]int{{0, 1, 0}, {0, 0, 0}, {0, 0, 0}}))
	t.Run("PlayerB cell 12 k=1", testApplyPlayerInput(1, 2, 1,
		[3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		[3][3]int{{0, 2, 0}, {0, 0, 0}, {0, 0, 0}}))
	t.Run("PlayerA cell 12 k=8", testApplyPlayerInput(1, 2, 8,
		[3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		[3][3]int{{0, 1, 0}, {0, 0, 0}, {0, 0, 0}}))
	t.Run("PlayerB cell 12 k=9", testApplyPlayerInput(1, 2, 9,
		[3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		[3][3]int{{0, 2, 0}, {0, 0, 0}, {0, 0, 0}}))
}

func testApplyPlayerInput(r, c, k int, gameState [3][3]int, expected [3][3]int) func(*testing.T) {
	return func(t *testing.T) {
		actual := applyPlayerInput(r, c, k, gameState)
		if actual != expected {
			t.Error(fmt.Sprintf("With entry %v%v and k %v expected %v got %v. ", r, c, k, expected, actual))
		}
	}
}

func TestGetWinner(t *testing.T) {
	t.Run("Second row", testGetWinner(
		[3][3]int{{2, 1, 2}, {1, 1, 1}, {0, 0, 2}}, 6, 1))
	t.Run("Top left diagonal", testGetWinner(
		[3][3]int{{2, 1, 2}, {1, 2, 1}, {1, 0, 2}}, 7, 2))
}

func testGetWinner(gameState [3][3]int, k int, expected int) func(*testing.T) {
	return func(t *testing.T) {
		actual := getWinner(gameState, k)
		if actual != expected {
			t.Error(fmt.Sprintf("With %v expected %v got %v. ", gameState, expected, actual))
		}
	}
}

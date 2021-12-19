package main

import (
	"testing"
)

func TestHasVerticalBingo(t *testing.T) {
	card := &BingoCard{}

	card.marked[0][0] = true
	card.marked[1][0] = true
	card.marked[2][0] = true
	card.marked[3][0] = true
	card.marked[4][0] = true

	if !card.HasVerticalBingo() {
		t.Error("First vertical Bingo not detected")
	}

	card.marked[0][0] = false

	if card.HasVerticalBingo() {
		t.Error("False positive")
	}

	card.marked[0][2] = true
	card.marked[1][2] = true
	card.marked[2][2] = true
	card.marked[3][2] = true
	card.marked[4][2] = true

	if !card.HasVerticalBingo() {
		t.Error("Second vertical Bingo not detected")
	}

	card.marked[0][2] = false

	if card.HasVerticalBingo() {
		t.Error("False positive")
	}

	card = &BingoCard{}
	card.marked[0][0] = true
	card.marked[0][1] = true
	card.marked[0][2] = true
	card.marked[0][3] = true
	card.marked[0][4] = true

	if card.HasVerticalBingo() {
		t.Error("False positive")
	}

}

func TestHasHorizontalBingo(t *testing.T) {
	card := &BingoCard{}

	card.marked[0][0] = true
	card.marked[0][1] = true
	card.marked[0][2] = true
	card.marked[0][3] = true
	card.marked[0][4] = true

	if !card.HasHorizontalBingo() {
		t.Error("First horizontal Bingo not detected")
	}

	card.marked[0][0] = false

	if card.HasHorizontalBingo() {
		t.Error("False positive")
	}

	card.marked[2][0] = true
	card.marked[2][1] = true
	card.marked[2][2] = true
	card.marked[2][3] = true
	card.marked[2][4] = true

	if !card.HasHorizontalBingo() {
		t.Error("Second horizontal Bingo not detected")
	}

	card.marked[2][0] = false

	if card.HasHorizontalBingo() {
		t.Error("False positive")
	}

	card = &BingoCard{}
	card.marked[0][0] = true
	card.marked[1][0] = true
	card.marked[2][0] = true
	card.marked[3][0] = true
	card.marked[4][0] = true

	if card.HasHorizontalBingo() {
		t.Error("False positive")
	}

}

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

func TestParseCalledNumbers(t *testing.T) {
	testStr := "7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1"

	numSlice := ParseCalledNumbers(testStr)

	if numSlice[0] != 7 {
		t.Error("didn't get 7")
	}

	if numSlice[3] != 5 {
		t.Error("didn't get 5")
	}
}

func TestParseBingoCardLine(t *testing.T) {
	testStr := " 6 10  3 18  5"

	numSlice := ParseBingoCardLine(testStr)

	if numSlice[0] != 6 {
		t.Error("first field not parsed")
	}

	if numSlice[4] != 5 {
		t.Error("last field not parsed")
	}
}

func TestParseDay04aInputFile(t *testing.T) {
	bingoCards := ParseDay04aInputFile("test_data_day04a.txt")

	t.Error(bingoCards)

	//fmt.Print(bingoCards)
}

package main

import (
	"strconv"
	"strings"
)

const BINGO_CARD_ROWS = 5
const BINGO_CARD_COLS = 5

type BingoCard struct {
	numbers [5][5]int
	marked  [5][5]bool
}

func (card *BingoCard) HasVerticalBingo() bool {

	maxRow := 0

	for i := 0; i < BINGO_CARD_ROWS; i++ {
		maxRow = 0
		for j := 0; j < BINGO_CARD_COLS; j++ {
			if card.marked[j][i] {
				maxRow++
			} else {
				break
			}
		}
		if maxRow == BINGO_CARD_ROWS {
			break
		}
	}

	return maxRow == BINGO_CARD_ROWS
}

func (card *BingoCard) HasHorizontalBingo() bool {

	maxColumn := 0

	for i := 0; i < BINGO_CARD_COLS; i++ {
		maxColumn = 0
		for j := 0; j < BINGO_CARD_ROWS; j++ {
			if card.marked[i][j] {
				maxColumn++
			} else {
				break
			}
		}
		if maxColumn == BINGO_CARD_COLS {
			break
		}
	}

	return maxColumn == BINGO_CARD_COLS
}

func ParseCalledNumbers(calledNumberLine string) (calledNumbers []int) {

	calledNumberStrings := strings.Split(calledNumberLine, ",")
	calledNumbers = make([]int, len(calledNumberStrings))

	for i, calledNumberString := range calledNumberStrings {
		n, _ := strconv.Atoi(calledNumberString)
		calledNumbers[i] = n
	}

	return calledNumbers
}

func ParseBingoCardLine(line string) (lineNumbers []int) {
	lineNumbers = make([]int, 5)

	theStrings := strings.Fields(line)

	for i, theNumber := range theStrings {
		n, _ := strconv.Atoi(theNumber)
		lineNumbers[i] = n
	}

	return lineNumbers
}

func ParseDay04aInputFile(filename string) ([]int, []*BingoCard) {
	stringLines := GetStringsFromFile(filename)

	bingoCards := make([]*BingoCard, 0)

	calledNumbers := ParseCalledNumbers(stringLines[0])

	cardCounter := 0
	bingoCard := &BingoCard{}
	for i := 1; i < len(stringLines); i++ {

		if stringLines[i] == "" {
			continue
		}

		numSlice := ParseBingoCardLine(stringLines[i])

		for n, val := range numSlice {
			bingoCard.numbers[cardCounter][n] = val
		}

		cardCounter++

		if cardCounter > 4 {
			bingoCards = append(bingoCards, bingoCard)
			cardCounter = 0
			bingoCard = &BingoCard{}
		}
	}

	return calledNumbers, bingoCards
}

func (card *BingoCard) MarkNumberIfPresent(theNumber int) {
LAST:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if card.numbers[i][j] == theNumber {
				card.marked[i][j] = true
				break LAST
			}
		}
	}
}

func PlayCards(calledNumbers []int, cards []*BingoCard) (returnCard *BingoCard, lastNumberCalled int) {
LAST:
	for _, number := range calledNumbers {
		for _, card := range cards {
			card.MarkNumberIfPresent(number)
			if card.HasHorizontalBingo() || card.HasVerticalBingo() {
				returnCard = card
				lastNumberCalled = number
				break LAST
			}
		}
	}

	return returnCard, lastNumberCalled
}

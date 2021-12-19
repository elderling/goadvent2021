package main

const BINGO_CARD_ROWS = 5
const BINGO_CARD_COLS = 5

type BingoCard struct {
	//numbers [BINGO_CARD_ROWS][BINGO_CARD_COLS]int
	marked [5][5]bool
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

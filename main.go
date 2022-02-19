package main

import (
	"fmt"
)

func main() {
	ints := ReadIntegersFromFile("my_day01a_data.txt")

	fmt.Printf("Day01a: Total Number of increases %v\n", CountTotalIncreases(ints))

	fmt.Printf("Day01b: Total Number of sliding increases %v\n", CountTotalIncreases((DepthsToSlidingDepths(ints))))

	finalLoc := &SubmarineLocation{
		totalDistance: 0,
		depth:         0,
	}

	fmt.Printf("Day02a: Solution: %v\n", finalLoc.Day02aSolution("my_day02a_data.txt"))

	finalLoc.totalDistance = 0
	finalLoc.depth = 0
	fmt.Printf("Day02b: Solution: %v\n", finalLoc.Day02bSolution("my_day02a_data.txt"))

	calledNumbers, bingoCards := ParseDay04aInputFile("my_day04a_data.txt")
	winnerCard, lastNumberCalled := PlayCards(calledNumbers, bingoCards)
	score := DoCardArithmetic(winnerCard, lastNumberCalled)
	fmt.Printf("Day04a: Solution: %v\n", score)
}

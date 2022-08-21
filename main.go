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

	score = DoCardArithmetic(FindLastWinningCard(ParseDay04aInputFile("my_day04a_data.txt")))

	fmt.Printf("Day04b: Solution: %v\n", score)

	bad_spots := Day05aSolution("my_day05a_data.txt")
	fmt.Printf("Day05a: Solution: %v\n", bad_spots)

	fish := ReadFishFromFile("my_day06a_data.txt")
	fish = FishDays(fish, 80)
	total_fish := len(fish)
	fmt.Printf("Day06a: Solution: %v\n", total_fish)

	fishb := ReadFishFromFile("my_day06a_data.txt")
	bucket := FishToFishBucket(fishb)
	fmt.Printf("Day06b: Solution: %v\n", SolutionDay06b(bucket, 256))

	fmt.Printf("Day07a: Solution: %v\n", SolutionDay07a(ReadCommaSeparatedFromFile("my_day07a_data.txt")))

}

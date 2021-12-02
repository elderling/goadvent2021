package main

import (
	"fmt"
)

func main() {
	ints := ReadIntegersFromFile("my_day01a_data.txt")

	fmt.Printf("Day01a: Total Number of increases %v\n", CountTotalIncreases(ints))

	fmt.Printf("Day01b: Total Number of sliding increases %v\n", CountTotalIncreases((DepthsToSlidingDepths(ints))))
}

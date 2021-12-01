package main

import (
	"fmt"
)

func main() {
	ints := ReadIntegersFromFile("my_day01a_data.txt")

	fmt.Printf("Day01a: Total Number of increases %v\n", CountTotalIncreases(ints))
}

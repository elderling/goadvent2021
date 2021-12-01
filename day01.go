package main

import (
	"bufio"
	"os"
	"strconv"
)

func CountTotalIncreases(depths []int) int {
	total := 0
	prev := 0
	for _, d := range depths {
		if prev > 0 {
			if d > prev {
				total++
			}
		}
		prev = d
	}

	return total
}

func ReadIntegersFromFile(filename string) []int {

	ints := []int{}

	f, err := os.Open(filename)
	if err != nil {
		panic("couldn't open file")
	}

	input := bufio.NewScanner(f)

	for input.Scan() {
		theInt, _ := strconv.Atoi(input.Text())

		ints = append(ints, theInt)
	}

	return ints
}

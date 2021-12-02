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

func GetSlidingWindow(ints []int, i int) (sum int, ok bool) {
	sum = 0
	ok = true

	if i > len(ints)-3 {
		return 0, false
	}

	for j := i; j < i+3; j++ {
		sum = sum + ints[j]
	}

	return sum, ok
}

func DepthsToSlidingDepths(ints []int) []int {
	slidings := []int{}
	for i := range ints {
		sum, ok := GetSlidingWindow(ints, i)
		if ok {
			slidings = append(slidings, sum)
		}
	}

	return slidings
}

package main

import (
	"math"
	"strconv"
)

func Day03aSolution(filename string) int {
	var ones [64]int
	var zeros [64]int
	s := GetStringsFromFile(filename)

	width := len(s[0])

	n := BinStringsToInts(s)

	for _, theInt := range n {
		for pos := 0; pos < width; pos++ {
			if BitAtPosition(theInt, pos) {
				ones[pos]++
			} else {
				zeros[pos]++
			}
		}

	}

	gamma := 0
	epsilon := 0
	for pos := 0; pos < width; pos++ {
		if ones[pos] > zeros[pos] {
			gamma |= int(math.Exp2(float64(pos)))
		} else {
			epsilon |= int(math.Exp2(float64(pos)))

		}
	}

	return gamma * epsilon
}

func ReverseString(s string) string {
	reversedBytes := make([]byte, 0)
	r := []byte(s)
	for i := len(r) - 1; i >= 0; i-- {
		reversedBytes = append(reversedBytes, r[i])
	}
	return string(reversedBytes)
}

func BitAtPosition(n int, pos int) bool {
	mask := int(math.Exp2(float64(pos)))
	return n&mask > 0
}

func BinStringsToInts(s []string) []int {
	r := make([]int, 0)

	for _, s := range s {
		theInt, _ := strconv.ParseInt(s, 2, 64)
		r = append(r, int(theInt))
	}

	return r
}

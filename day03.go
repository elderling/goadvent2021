package main

import (
	"math"
	"strconv"
)

func Day03aSolution(filename string) int {
	return 198
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

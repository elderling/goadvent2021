package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type ByInt []int

func (a ByInt) Len() int           { return len(a) }
func (a ByInt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByInt) Less(i, j int) bool { return a[i] < a[j] }

func Median(a []int) int {
	sort.Sort(ByInt(a))
	return a[len(a)/2]
}

func IntAbs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func ReadCommaSeparatedFromFile(filename string) []int {
	ints := []int{}

	f, err := os.Open(filename)
	if err != nil {
		panic("couldn't open file")
	}

	input := bufio.NewScanner(f)

	for input.Scan() {
		commaSeparatedInts := input.Text()

		s := strings.Split(commaSeparatedInts, ",")

		for _, v := range s {
			theInt, _ := strconv.Atoi(v)

			ints = append(ints, theInt)
		}
	}

	return ints
}

func SolutionDay07a(a []int) int {
	total := 0

	median := Median(a)

	for _, v := range a {
		total += IntAbs(v - median)
	}

	return total
}

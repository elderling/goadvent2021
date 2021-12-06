package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type SubmarineCommand struct {
	direction string
	distance  int
}

type SubmarineLocation struct {
	totalDistance int
	depth         int
	aim           int
}

func (sl *SubmarineLocation) Day02aSolution(filename string) int {
	commands := GetStringsFromFile(filename)

	for _, c := range commands {
		subLoc := ParseSubmarineCommand(c)

		sl.DoCommand(subLoc)
	}
	return sl.totalDistance * sl.depth
}

func (sl *SubmarineLocation) Day02bSolution(filename string) int {
	commands := GetStringsFromFile(filename)

	for _, c := range commands {
		subLoc := ParseSubmarineCommand(c)

		sl.DoCommand2b(subLoc)
	}
	return sl.totalDistance * sl.depth
}

func (sl *SubmarineLocation) DoCommand(sc *SubmarineCommand) {
	switch sc.direction {
	case "forward":
		sl.totalDistance = sl.totalDistance + sc.distance
	case "up":
		sl.depth = sl.depth - sc.distance
	case "down":
		sl.depth = sl.depth + sc.distance
	}
}

func (sl *SubmarineLocation) DoCommand2b(sc *SubmarineCommand) {
	switch sc.direction {
	case "forward":
		{
			sl.totalDistance = sl.totalDistance + sc.distance
			sl.depth = sl.depth + sc.distance*sl.aim
		}
	case "up":
		sl.aim = sl.aim - sc.distance
	case "down":
		sl.aim = sl.aim + sc.distance
	}
}

func GetStringsFromFile(filename string) []string {

	vals := []string{}

	f, err := os.Open(filename)
	if err != nil {
		panic("couldn't open file")
	}

	input := bufio.NewScanner(f)

	for input.Scan() {
		theVal := input.Text()

		vals = append(vals, theVal)
	}

	return vals
}

func ParseSubmarineCommand(c string) *SubmarineCommand {
	sc := new(SubmarineCommand)

	tokens := strings.Split(c, " ")

	sc.direction = tokens[0]
	sc.distance, _ = strconv.Atoi(tokens[1])

	return sc
}

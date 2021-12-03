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
}

func (sl *SubmarineLocation) Day02aSolution() int {
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

func GetCommandsFromFile(filename string) []string {

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

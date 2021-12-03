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

func Day02aSolution() int {
	return 0
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

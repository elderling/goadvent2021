package main

import (
	"bufio"
	"os"
)

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

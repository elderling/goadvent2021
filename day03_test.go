package main

import (
	"testing"
)

func TestDay03aSolution(t *testing.T) {
	solution := Day03aSolution("test_data_day03a.txt")
	if solution != 198 {
		t.Error("bad TestDay03aSolution")
	}
}

package main

import (
	"testing"
)

func TestCountTotalIncreases(t *testing.T) {
	day01TestInput := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}

	total := CountTotalIncreases(day01TestInput)

	if total != 7 {
		t.Errorf("Got incorrect number of total expected 7 got %v", total)
	}
}

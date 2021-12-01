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

func TestReadIntegersFromFile(t *testing.T) {
	integers := ReadIntegersFromFile("test_data_day01a.txt")

	if len(integers) != 3 {
		t.Errorf("Got incorrect number of ints. Expected 3 got %v", len(integers))
	}
}

func TestGetSlidingWindow(t *testing.T) {
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

	sum, ok := GetSlidingWindow(day01TestInput, 2)
	if !ok {
		t.Error("Didn't get expected ok")
	}
	if sum != 208+210+200 {
		t.Error("didn't get correct window sum")
	}
	_, ok = GetSlidingWindow(day01TestInput, 9)

	if ok != false {
		t.Error("Bounds checking broken")
	}
}

package main

import "testing"

func TestMedian(t *testing.T) {
	n := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	if Median(n) != 2 {
		t.Error("Median not being calculated correctly")
	}
}

func TestSolutionDay07a(t *testing.T) {
	n := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	if SolutionDay07a(n) != 37 {
		t.Error("TestSolutionDay07a is broken")
	}
}

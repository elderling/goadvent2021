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

func TestReverseString(t *testing.T) {

	r := ReverseString("")
	if r != "" {
		t.Errorf("Failed to reverse empty string. Got '%v'", r)
	}

	r = ReverseString("a")
	if r != "a" {
		t.Errorf("Failed to reverse 'a' Got: '%v'", r)
	}

	r = ReverseString("ab")
	if r != "ba" {
		t.Errorf("Failed to reverse 'ab' Got: '%v'", r)
	}

	r = ReverseString("racecar")
	if r != "racecar" {
		t.Errorf("Failed to reverse 'racecar' Got: '%v'", r)
	}

	r = ReverseString("abcdef")
	if r != "fedcba" {
		t.Errorf("Failed to reverse 'abcdef' Got: '%v'", r)
	}
}

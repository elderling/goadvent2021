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

func TestBinStringsToInts(t *testing.T) {
	s := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	r := BinStringsToInts(s)

	if r[0] != 4 {
		t.Error("Conversion Error")
	}
}

func TestBitAtPosition(t *testing.T) {
	if !BitAtPosition(4, 2) {
		t.Error("4,2 busted")
	}
	if !BitAtPosition(8, 3) {
		t.Error("8,3 busted")
	}
	if !BitAtPosition(8+4, 2) {
		t.Error("8+4,2 busted")
	}
}

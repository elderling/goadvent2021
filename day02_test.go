package main

import (
	"testing"
)

func TestDay02aSolution(t *testing.T) {
	if Day02aSolution() != 150 {
		t.Errorf("Wrong solution. Expected 150, got %v\n", Day02aSolution())
	}

}

func TestGetCommandsFromFile(t *testing.T) {
	expected := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	commands := GetCommandsFromFile("test_data_day02a.txt")

	for i, val := range expected {
		if commands[i] != val {
			t.Errorf("Bad GetCommandsFromFile. Expected: %v, got %v", val, commands[i])
		}
	}

}

func TestParseSubmarineCommand(t *testing.T) {
	subCommand := ParseSubmarineCommand("forward 5")

	if subCommand.direction != "forward" {
		t.Errorf("ParseSubmarineCommand parse error, direction: Expected %v, got %v", "forward", subCommand.direction)
	}

	if subCommand.distance != 5 {
		t.Errorf("ParseSubmarineCommand parse error, direction: Expected %v, got %v", 5, subCommand.direction)
	}
}

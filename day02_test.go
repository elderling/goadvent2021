package main

import (
	"testing"
)

func TestDay02aSolution(t *testing.T) {
	finalLoc := &SubmarineLocation{
		totalDistance: 0,
		depth:         0,
	}
	if finalLoc.Day02aSolution("test_data_day02a.txt") != 150 {
		t.Errorf("Wrong solution. Expected 150, got %v\n", finalLoc.Day02aSolution("test_data_day02a.txt"))
	}

}

func TestGetStringsFromFile(t *testing.T) {
	expected := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	commands := GetStringsFromFile("test_data_day02a.txt")

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

func TestDoCommand(t *testing.T) {
	command := &SubmarineCommand{
		direction: "forward",
		distance:  10,
	}

	loc := &SubmarineLocation{
		totalDistance: 0,
		depth:         0,
	}

	loc.DoCommand(command)

	if loc.depth != 0 {
		t.Errorf("bad depth")
	}

	if loc.totalDistance != 10 {
		t.Errorf("bad distance")
	}

	command = &SubmarineCommand{
		direction: "down",
		distance:  1,
	}

	loc.DoCommand(command)

	if loc.depth != 1 {
		t.Errorf("bad depth")
	}

	if loc.totalDistance != 10 {
		t.Errorf("bad distance")
	}

	command = &SubmarineCommand{
		direction: "up",
		distance:  1,
	}

	loc.DoCommand(command)

	if loc.depth != 0 {
		t.Errorf("bad depth")
	}

	if loc.totalDistance != 10 {
		t.Errorf("bad distance")
	}

}

func TestDoCommand2b(t *testing.T) {
	theStrings := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	location := &SubmarineLocation{}

	for i, v := range theStrings {
		command := ParseSubmarineCommand(v)

		location.DoCommand2b(command)

		switch i {
		case 0: // forward 5
			if location.totalDistance != 5 && location.depth != 0 && location.aim != 0 {
				t.Errorf("Command '%v' not executed properly", v)
			}
		case 1: // down 5
			if location.totalDistance != 5 && location.depth != 0 && location.aim != 5 {
				t.Errorf("Command '%v' not executed properly", v)
			}
		case 2: // forward 8
			if location.totalDistance != 13 && location.depth != 40 && location.aim != 5 {
				t.Errorf("Command '%v' not executed properly", v)
			}
		case 3: // up 3
			if location.totalDistance != 13 && location.depth != 40 && location.aim != 2 {
				t.Errorf("Command '%v' not executed properly", v)
			}
		case 4: // down 8
			if location.totalDistance != 13 && location.depth != 40 && location.aim != 10 {
				t.Errorf("Command '%v' not executed properly", v)
			}
		case 5: // forward 2
			if location.totalDistance != 15 && location.depth != 60 && location.aim != 10 {
				t.Errorf("Command '%v' not executed properly", v)
			}
		}
	}

}

func TestDay2bSolution(t *testing.T) {
	sl := &SubmarineLocation{}

	sol := sl.Day02bSolution("test_data_day02a.txt")

	if sol != 900 {
		t.Errorf("Invalid solution. Expected 900, got: %v", sol)
	}
}

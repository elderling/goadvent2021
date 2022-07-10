package main

import "testing"

func TestStringPointToPoint(t *testing.T) {
	s := "0,0"
	point := StringPointToPoint(s)

	if point.x != 0 {
		t.Error("didn't get x == 0")
	}

	if point.y != 0 {
		t.Error("didn't get y == 0")
	}

	r := "1,1"

	point1 := StringPointToPoint(r)

	if point1.x != 1 {
		t.Error("didn't get x == 1")
	}

	if point1.y != 1 {
		t.Error("didn't get y == 1")
	}
}

func TestStringToLine(t *testing.T) {
	s := "0,0 -> 1,1"

	line := StringToLine(s)

	if line.i.x != 0 {
		t.Error("problem with line.i.x")
	}
	if line.i.y != 0 {
		t.Error("problem with line.i.y")
	}
	if line.j.x != 1 {
		t.Error("problem with line.j.x")
	}
	if line.j.y != 1 {
		t.Error("problem with line.j.y")
	}
}

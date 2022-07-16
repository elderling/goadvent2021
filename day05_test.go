package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

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

func TestGepPointsInLine(t *testing.T) {
	l1 := StringToLine("0,9 -> 2,9")

	p1 := Point{x: 0, y: 9}
	p2 := Point{x: 1, y: 9}
	p3 := Point{x: 2, y: 9}

	gpil1 := GetPointsInLine(l1)
	if !slices.Equal(gpil1, []Point{p1, p2, p3}) {
		t.Error("not the right points")
	}

	l2 := StringToLine("2,9 -> 0,9")

	gpil2 := GetPointsInLine(l2)
	if !slices.Equal(gpil2, []Point{p1, p2, p3}) {
		t.Error("not the right points")
	}

	p4 := Point{x: 9, y: 2}
	p5 := Point{x: 9, y: 3}
	p6 := Point{x: 9, y: 4}

	gpil3 := GetPointsInLine(StringToLine("9,2 -> 9,4"))
	if !slices.Equal(gpil3, []Point{p4, p5, p6}) {
		t.Error("not the right points")
	}

	gpil4 := GetPointsInLine(StringToLine("9,4 -> 9,2"))
	if !slices.Equal(gpil4, []Point{p4, p5, p6}) {
		t.Error("not the right points")
	}
}

func TestParseDay05aInputFile(t *testing.T) {
	lines := ParsDay05aInputFile("test_data_day05a.txt")

	if lines[0].i.x != 0 {
		t.Error("read wrong")
	}

}

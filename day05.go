package main

import (
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Line struct {
	i *Point
	j *Point
}

func StringToLine(s string) *Line {
	stringPoints := strings.Split(s, " -> ")

	theLine := &Line{
		i: StringPointToPoint(stringPoints[0]),
		j: StringPointToPoint(stringPoints[1])}

	return theLine
}

func StringPointToPoint(s string) *Point {
	stringInts := strings.Split(s, ",")

	x, _ := strconv.Atoi(stringInts[0])
	y, _ := strconv.Atoi(stringInts[1])

	return &Point{x: x, y: y}
}

/*
func MarkLineOfVents(v *map[Point]int, l *Line) *map[Point]int {

	return v
}
*/

func GetPointsInLine(l *Line) []Point {

	p := make([]Point, 0)

	// x coordinate the same
	if l.i.x == l.j.x {
		if l.i.y <= l.j.y { //  count from i to j
			for q := l.i.y; q <= l.j.y; q++ {
				p = append(p, Point{y: q, x: l.i.x})
			}
		} else {
			for q := l.j.y; q <= l.i.y; q++ { // count from j to i
				p = append(p, Point{y: q, x: l.i.x})
			}
		}
	} else { // y coordinate the same
		if l.i.x <= l.j.x { // count from i to j
			for q := l.i.x; q <= l.j.x; q++ {
				p = append(p, Point{x: q, y: l.i.y})
			}
		} else { // count from j to i
			for q := l.j.x; q <= l.i.x; q++ {
				p = append(p, Point{x: q, y: l.i.y})
			}
		}
	}

	return p
}

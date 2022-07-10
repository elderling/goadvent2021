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

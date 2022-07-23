package main

import (
	"strconv"
	"strings"
)

func ReadFishFromFile(filename string) []int {

	lines := GetStringsFromFile(filename)

	line := lines[0]

	string_ints := strings.Split(line, ",")

	fish := make([]int, 0)
	for _, val := range string_ints {
		the_int, _ := strconv.Atoi(val)
		fish = append(fish, the_int)
	}

	return fish
}

func FishDay(f []int) []int {

	add_num_fish := 0
	for i, val := range f {
		val--
		if val < 0 {
			val = 6
			add_num_fish++
		}
		f[i] = val
	}

	for i := 0; i < add_num_fish; i++ {
		f = append(f, 8)
	}

	return f
}

func FishDays(f []int, days int) []int {
	for i := 0; i < days; i++ {
		f = FishDay(f)
	}
	return f
}

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

func FishToFishBucket(fish []int) map[int]int {
	bucket := make(map[int]int)

	for _, val := range fish {
		bucket[val]++
	}

	return bucket
}

func CountFishInBucket(bucket map[int]int) int {
	total := 0

	for i := 0; i <= 8; i++ {
		total += bucket[i]
	}

	return total
}

func BucketDay(bucket map[int]int) map[int]int {

	new_bucket := make(map[int]int)

	new_bucket[8] = bucket[0]
	new_bucket[7] = bucket[8]
	new_bucket[6] = bucket[0] + bucket[7]
	new_bucket[5] = bucket[6]
	new_bucket[4] = bucket[5]
	new_bucket[3] = bucket[4]
	new_bucket[2] = bucket[3]
	new_bucket[1] = bucket[2]
	new_bucket[0] = bucket[1]

	return new_bucket
}

func BucketDays(bucket map[int]int, days int) map[int]int {
	return_bucket := bucket

	for i := 0; i < days; i++ {
		return_bucket = BucketDay(return_bucket)
	}

	return return_bucket
}

func SolutionDay06b(bucket map[int]int, days int) int {
	return CountFishInBucket(BucketDays(bucket, days))
}

package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestReadFishFromFile(t *testing.T) {
	fish := ReadFishFromFile("test_data_day06a.txt")

	test_data := []int{3, 4, 3, 1, 2}

	if !slices.Equal(fish, test_data) {
		t.Error("Failed to read fish test data")
	}
}

func TestFishDay(t *testing.T) {
	test_data := []int{3, 4, 3, 1, 2}

	test_data = FishDay(test_data)

	if !slices.Equal(test_data, []int{2, 3, 2, 0, 1}) {
		t.Error("Day one bad")
	}

	test_data = FishDay(test_data)

	if !slices.Equal(test_data, []int{1, 2, 1, 6, 0, 8}) {
		t.Error("Day two bad")
	}
}

func TestFishDays(t *testing.T) {
	test_data := []int{3, 4, 3, 1, 2}

	eighteen_days := []int{6, 0, 6, 4, 5, 6, 0, 1, 1, 2, 6, 0, 1, 1, 1, 2, 2, 3, 3, 4, 6, 7, 8, 8, 8, 8}

	if !slices.Equal(FishDays(test_data, 18), eighteen_days) {
		t.Error("Something's broken with FishDays")
	}
}

func TestFishToFishBucket(t *testing.T) {
	bucket := FishToFishBucket([]int{3, 4, 3, 2})

	if bucket[3] != 2 {
		t.Error("bucket 3 broken")
	}
}

func TestCountFishInBucket(t *testing.T) {
	bucket := FishToFishBucket([]int{3, 4, 3, 2})
	total := CountFishInBucket(bucket)

	if total != 4 {
		t.Error("wrong number of fish")
	}

}

func TestSolutionDay06b(t *testing.T) {
	bucket := FishToFishBucket([]int{3, 4, 3, 1, 2})

	if SolutionDay06b(bucket, 18) != 26 {
		t.Error("Nope on SolutionDay06b")
	}

	if SolutionDay06b(bucket, 80) != 5934 {
		t.Error("Nope on SolutionDay06b")
	}
}

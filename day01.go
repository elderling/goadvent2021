package main

func CountTotalIncreases(depths []int) int {
	total := 0
	prev := 0
	for _, d := range depths {
		if prev > 0 {
			if d > prev {
				total++
			}
		}
		prev = d
	}

	return total
}

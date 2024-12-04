package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/Trashed/aoc2024/file"
)

func main() {

	// ========== Part one ============
	col1, col2 := file.ReadPuzzle("input.txt")
	fmt.Printf("Total distance is: %d\n", calcTotalDistance(col1, col2))

	// ========== Part two ============
	fmt.Printf("Total similarity is: %d\n", calcTotalSimilarity(col1, col2))
}

func calcTotalDistance(inputs1, inputs2 []int) int {
	distance := 0

	if len(inputs1) != len(inputs2) {
		panic("input array lengths doesn't match")
	}

	sort.Ints(inputs1)
	sort.Ints(inputs2)

	for i := 0; i < len(inputs1); i++ {
		d := math.Abs(float64(inputs1[i] - inputs2[i])) // We can take an absolute value since, for example, 3 - 2 = 1, and in reverse order 2 - 3 = -1
		distance += int(d)
	}

	return distance
}

func calcTotalSimilarity(inputs1, inputs2 []int) int {

	totalSimilarity := 0

	countOccurrance := func(v int, target []int) int {
		occurrence := 0
		for _, t := range target {
			if v == t {
				occurrence++
			}
		}
		return v * occurrence
	}

	for _, v := range inputs1 {
		totalSimilarity += countOccurrance(v, inputs2)
	}

	return totalSimilarity
}

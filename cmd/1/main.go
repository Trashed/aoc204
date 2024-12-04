package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/Trashed/aoc2024/file"
)

func main() {

	// ========== Part one ============

	distanceTotal := calcTotalDistance(file.ReadPuzzle("input.txt"))

	fmt.Printf("Total distance is: %d\n", distanceTotal)
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

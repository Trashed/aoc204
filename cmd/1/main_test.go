package main

import "testing"

func TestCalcTotalSimilarity(t *testing.T) {

	const expectedSimilarity int = 31

	vals1 := []int{3, 4, 2, 1, 3, 3}
	vals2 := []int{4, 3, 5, 3, 9, 3}

	actualSimilarity := calcTotalSimilarity(vals1, vals2)

	if expectedSimilarity != actualSimilarity {
		t.Fatalf("expected total similarity of %d but got %d", expectedSimilarity, actualSimilarity)
	}
}

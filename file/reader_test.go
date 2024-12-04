package file_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/Trashed/aoc2024/file"
)

func TestReadPuzzle(t *testing.T) {

	fileName := "test.txt"
	content := `3   4
4   3
2   5
1   3
3   9
3   3`

	if err := os.WriteFile(fileName, []byte(content), os.ModePerm); err != nil {
		t.Fatalf("failed to write test values to file: %s\n", err.Error())
	}

	// Test clean up
	defer func() {
		err := os.Remove(fileName)
		if err != nil {
			t.Fatalf("couldn't clean up test file: %s\n", err.Error())
		}
	}()

	vals1, vals2 := file.ReadPuzzle("test.txt")
	if vals1 == nil && vals2 == nil {
		t.Fatalf("values read from the file shouldn't result to nil")
	}

	if len(vals1) == 0 || len(vals2) == 0 {
		t.Fatal("length of the slices shouldn't be 0")
	}

	expectedVals1 := []int{3, 4, 2, 1, 3, 3}
	expectedVals2 := []int{4, 3, 5, 3, 9, 3}

	if ok, err := assertEquals(vals1, expectedVals1); !ok {
		t.Fatal(err.Error())
	}

	if ok, err := assertEquals(vals2, expectedVals2); !ok {
		t.Fatal(err.Error())
	}
}

func assertEquals(actual, expected []int) (bool, error) {
	result := false
	if len(actual) != len(expected) {
		return result, fmt.Errorf("slice lengths don't match: expected %d but got %d", len(expected), len(actual))
	}

	for i, val := range actual {
		result = val == expected[i]
		if !result {
			return result, fmt.Errorf("slice values on index %d don't match: expected %d but got %d", i, expected[i], val)
		}
	}

	return result, nil
}

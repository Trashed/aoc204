package main

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	file, err := os.OpenFile("input.txt", os.O_RDONLY, fs.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	distanceTotal := calcTotalDistance(readColumns(file))

	fmt.Printf("Total distance is: %d\n", distanceTotal)
}

func readColumns(r io.Reader) ([]int, []int) {
	col1, col2 := []int{}, []int{}

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()

		vals := strings.Split(line, "   ")
		i1, _ := strconv.ParseInt(vals[0], 10, 64)
		i2, _ := strconv.ParseInt(vals[1], 10, 64)

		col1 = append(col1, int(i1))
		col2 = append(col2, int(i2))

	}

	return col1, col2
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
